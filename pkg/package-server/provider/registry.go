package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/operator-framework/operator-registry/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/util/labels"

	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/informers/externalversions"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/queueinformer"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators"
)

const (
	defaultConnectionTimeout = 5 * time.Second
)

type sourceKey struct {
	name      string
	namespace string
}

type registryClient struct {
	api.RegistryClient
	source *operatorsv1alpha1.CatalogSource
	conn   *grpc.ClientConn
}

func newRegistryClient(source *operatorsv1alpha1.CatalogSource, conn *grpc.ClientConn) registryClient {
	return registryClient{
		RegistryClient: api.NewRegistryClient(conn),
		source:         source,
		conn:           conn,
	}
}

// RegistryProvider aggregates several `CatalogSources` and establishes gRPC connections to their registry servers.
type RegistryProvider struct {
	queueinformer.Operator

	mu              sync.RWMutex
	globalNamespace string
	clients         map[sourceKey]registryClient
}

var _ PackageManifestProvider = &RegistryProvider{}

func NewRegistryProvider(ctx context.Context, crClient versioned.Interface, operator queueinformer.Operator, wakeupInterval time.Duration, watchedNamespaces []string, globalNamespace string) (*RegistryProvider, error) {
	p := &RegistryProvider{
		Operator: operator,

		globalNamespace: globalNamespace,
		clients:         make(map[sourceKey]registryClient),
	}

	for _, namespace := range watchedNamespaces {
		informerFactory := externalversions.NewSharedInformerFactoryWithOptions(crClient, wakeupInterval, externalversions.WithNamespace(namespace))
		catsrcInformer := informerFactory.Operators().V1alpha1().CatalogSources()

		// Register queue and QueueInformer
		logrus.WithField("namespace", namespace).Info("watching catalogsources")
		catsrcQueueInformer, err := queueinformer.NewQueueInformer(
			ctx,
			queueinformer.WithInformer(catsrcInformer.Informer()),
			queueinformer.WithSyncer(queueinformer.LegacySyncHandler(p.syncCatalogSource).ToSyncerWithDelete(p.catalogSourceDeleted)),
		)
		if err != nil {
			return nil, err
		}
		p.RegisterQueueInformer(catsrcQueueInformer)
	}

	return p, nil
}

func (p *RegistryProvider) getClient(key sourceKey) (registryClient, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	client, ok := p.clients[key]
	return client, ok
}

func (p *RegistryProvider) setClient(client registryClient, key sourceKey) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.clients[key] = client
}

func (p *RegistryProvider) removeClient(key sourceKey) (registryClient, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	client, ok := p.clients[key]
	if !ok {
		return registryClient{}, false
	}

	delete(p.clients, key)
	return client, true
}

func (p *RegistryProvider) syncCatalogSource(obj interface{}) (syncError error) {
	source, ok := obj.(*operatorsv1alpha1.CatalogSource)
	if !ok {
		logrus.Errorf("catalogsource type assertion failed: wrong type: %#v", obj)
	}

	logger := logrus.WithFields(logrus.Fields{
		"action":    "sync catalogsource",
		"name":      source.GetName(),
		"namespace": source.GetNamespace(),
	})

	if source.Status.RegistryServiceStatus == nil {
		logger.Debug("registry service is not ready for grpc connection")
		return
	}

	key := sourceKey{source.GetName(), source.GetNamespace()}
	client, ok := p.getClient(key)
	if ok && source.Status.RegistryServiceStatus.ServiceName != "" {
		logger.Info("update detected, attempting to reset grpc connection")
		client.conn.ResetConnectBackoff()

		ctx, cancel := context.WithTimeout(context.TODO(), defaultConnectionTimeout)
		defer cancel()

		changed := client.conn.WaitForStateChange(ctx, connectivity.TransientFailure)
		if !changed {
			logger.Debugf("grpc connection reset timeout")
			syncError = fmt.Errorf("grpc connection reset timeout")
			return
		}

		logger.Info("grpc connection reset")
		return
	} else if ok {
		// Address type grpc CatalogSource, drop the connection dial in to the new address
		client.conn.Close()
	}

	logger.Info("attempting to add a new grpc connection")
	conn, err := grpc.Dial(source.Address(), grpc.WithInsecure())
	if err != nil {
		logger.WithField("err", err.Error()).Errorf("could not connect to registry service")
		syncError = err
		return
	}

	p.setClient(newRegistryClient(source, conn), key)
	logger.Info("new grpc connection added")

	return
}

func (p *RegistryProvider) catalogSourceDeleted(obj interface{}) {
	catsrc, ok := obj.(metav1.Object)
	if !ok {
		if !ok {
			tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %#v", obj))
				return
			}

			catsrc, ok = tombstone.Obj.(metav1.Object)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a Namespace %#v", obj))
				return
			}
		}
	}

	logger := logrus.WithFields(logrus.Fields{
		"action":    "CatalogSource Deleted",
		"name":      catsrc.GetName(),
		"namespace": catsrc.GetNamespace(),
	})
	logger.Debugf("attempting to remove grpc connection")

	key := sourceKey{catsrc.GetName(), catsrc.GetNamespace()}
	client, removed := p.removeClient(key)
	if removed {
		err := client.conn.Close()
		if err != nil {
			logger.WithField("err", err.Error()).Error("error closing connection")
			utilruntime.HandleError(fmt.Errorf("error closing connection %s", err.Error()))
			return
		}
		logger.Debug("grpc connection removed")
		return
	}

	logger.Debugf("no gRPC connection to remove")

}

func (p *RegistryProvider) Get(namespace, name string) (*operators.PackageManifest, error) {
	logger := logrus.WithFields(logrus.Fields{
		"action":    "Get PackageManifest",
		"name":      name,
		"namespace": namespace,
	})

	pkgs, err := p.List(namespace)
	if err != nil {
		return nil, fmt.Errorf("could not list packages in namespace %s", namespace)
	}

	for _, pkg := range pkgs.Items {
		if pkg.GetName() == name {
			return &pkg, nil
		}
	}

	logger.Info("package not found")
	return nil, nil
}

func (p *RegistryProvider) List(namespace string) (*operators.PackageManifestList, error) {
	logger := logrus.WithFields(logrus.Fields{
		"action":    "List PackageManifests",
		"namespace": namespace,
	})

	p.mu.RLock()
	defer p.mu.RUnlock()

	pkgs := []operators.PackageManifest{}
	for _, client := range p.clients {
		if client.source.GetNamespace() == namespace || client.source.GetNamespace() == p.globalNamespace || namespace == metav1.NamespaceAll {
			logger.Debugf("found CatalogSource %s", client.source.GetName())

			stream, err := client.ListPackages(context.Background(), &api.ListPackageRequest{})
			if err != nil {
				logger.WithField("err", err.Error()).Warnf("error getting stream")
				continue
			}
			for {
				pkgName, err := stream.Recv()
				if err == io.EOF {
					break
				}

				if err != nil {
					logger.WithField("err", err.Error()).Warnf("error getting data")
					break
				}
				pkg, err := client.GetPackage(context.Background(), &api.GetPackageRequest{Name: pkgName.GetName()})
				if err != nil {
					logger.WithField("err", err.Error()).Warnf("error getting package")
					continue
				}
				newPkg, err := toPackageManifest(logger, pkg, client)
				if err != nil {
					logger.WithField("err", err.Error()).Warnf("eliding package: error converting to packagemanifest")
					continue
				}

				// Set request namespace to stop kube clients from complaining about global namespace mismatch.
				if namespace != metav1.NamespaceAll {
					newPkg.SetNamespace(namespace)
				}
				pkgs = append(pkgs, *newPkg)
			}
		}
	}

	return &operators.PackageManifestList{Items: pkgs}, nil
}

func toPackageManifest(logger *logrus.Entry, pkg *api.Package, client registryClient) (*operators.PackageManifest, error) {
	pkgChannels := pkg.GetChannels()
	catsrc := client.source
	manifest := &operators.PackageManifest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pkg.GetName(),
			Namespace: catsrc.GetNamespace(),
			Labels: labels.CloneAndAddLabel(
				labels.CloneAndAddLabel(catsrc.GetLabels(),
					"catalog", catsrc.GetName()), "catalog-namespace", catsrc.GetNamespace()),
			CreationTimestamp: catsrc.GetCreationTimestamp(),
		},
		Status: operators.PackageManifestStatus{
			CatalogSource:            catsrc.GetName(),
			CatalogSourceDisplayName: catsrc.Spec.DisplayName,
			CatalogSourcePublisher:   catsrc.Spec.Publisher,
			CatalogSourceNamespace:   catsrc.GetNamespace(),
			PackageName:              pkg.Name,
			DefaultChannel:           pkg.GetDefaultChannelName(),
		},
	}

	var (
		providerSet   bool
		defaultElided bool
	)
	for _, pkgChannel := range pkgChannels {
		bundle, err := client.GetBundleForChannel(context.Background(), &api.GetBundleInChannelRequest{PkgName: pkg.GetName(), ChannelName: pkgChannel.GetName()})
		if err != nil {
			logger.WithError(err).WithField("channel", pkgChannel.GetName()).Warn("error getting bundle, eliding channel")
			defaultElided = defaultElided || pkgChannel.Name == manifest.Status.DefaultChannel
			continue
		}

		csv := operatorsv1alpha1.ClusterServiceVersion{}
		err = json.Unmarshal([]byte(bundle.GetCsvJson()), &csv)
		if err != nil {
			logger.WithError(err).WithField("channel", pkgChannel.GetName()).Warn("error unmarshaling csv, eliding channel")
			defaultElided = defaultElided || pkgChannel.Name == manifest.Status.DefaultChannel
			continue
		}
		manifest.Status.Channels = append(manifest.Status.Channels, operators.PackageChannel{
			Name:           pkgChannel.GetName(),
			CurrentCSV:     csv.GetName(),
			CurrentCSVDesc: operators.CreateCSVDescription(&csv),
		})

		if manifest.Status.DefaultChannel != "" && pkgChannel.GetName() == manifest.Status.DefaultChannel || !providerSet {
			manifest.Status.Provider = operators.AppLink{
				Name: csv.Spec.Provider.Name,
				URL:  csv.Spec.Provider.URL,
			}
			manifest.ObjectMeta.Labels["provider"] = manifest.Status.Provider.Name
			manifest.ObjectMeta.Labels["provider-url"] = manifest.Status.Provider.URL
			providerSet = true
		}
	}

	if len(manifest.Status.Channels) == 0 {
		return nil, fmt.Errorf("packagemanifest has no valid channels")
	}

	if defaultElided {
		logger.Warn("default channel elided, setting as first in packagemanifest")
		manifest.Status.DefaultChannel = manifest.Status.Channels[0].Name
	}

	return manifest, nil
}
