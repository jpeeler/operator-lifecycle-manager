// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceReference) DeepCopyInto(out *APIResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceReference.
func (in *APIResourceReference) DeepCopy() *APIResourceReference {
	if in == nil {
		return nil
	}
	out := new(APIResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceDefinitions) DeepCopyInto(out *APIServiceDefinitions) {
	*out = *in
	if in.Owned != nil {
		in, out := &in.Owned, &out.Owned
		*out = make([]APIServiceDescription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]APIServiceDescription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceDefinitions.
func (in *APIServiceDefinitions) DeepCopy() *APIServiceDefinitions {
	if in == nil {
		return nil
	}
	out := new(APIServiceDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceDescription) DeepCopyInto(out *APIServiceDescription) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]APIResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.StatusDescriptors != nil {
		in, out := &in.StatusDescriptors, &out.StatusDescriptors
		*out = make([]StatusDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpecDescriptors != nil {
		in, out := &in.SpecDescriptors, &out.SpecDescriptors
		*out = make([]SpecDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ActionDescriptor != nil {
		in, out := &in.ActionDescriptor, &out.ActionDescriptor
		*out = make([]ActionDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceDescription.
func (in *APIServiceDescription) DeepCopy() *APIServiceDescription {
	if in == nil {
		return nil
	}
	out := new(APIServiceDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionDescriptor) DeepCopyInto(out *ActionDescriptor) {
	*out = *in
	if in.XDescriptors != nil {
		in, out := &in.XDescriptors, &out.XDescriptors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(json.RawMessage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionDescriptor.
func (in *ActionDescriptor) DeepCopy() *ActionDescriptor {
	if in == nil {
		return nil
	}
	out := new(ActionDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppLink) DeepCopyInto(out *AppLink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppLink.
func (in *AppLink) DeepCopy() *AppLink {
	if in == nil {
		return nil
	}
	out := new(AppLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleJob) DeepCopyInto(out *BundleJob) {
	*out = *in
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleJob.
func (in *BundleJob) DeepCopy() *BundleJob {
	if in == nil {
		return nil
	}
	out := new(BundleJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleLookup) DeepCopyInto(out *BundleLookup) {
	*out = *in
	if in.BundleJob != nil {
		in, out := &in.BundleJob, &out.BundleJob
		*out = new(BundleJob)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ConfigMapResourceReference)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleLookup.
func (in *BundleLookup) DeepCopy() *BundleLookup {
	if in == nil {
		return nil
	}
	out := new(BundleLookup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDDescription) DeepCopyInto(out *CRDDescription) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]APIResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.StatusDescriptors != nil {
		in, out := &in.StatusDescriptors, &out.StatusDescriptors
		*out = make([]StatusDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpecDescriptors != nil {
		in, out := &in.SpecDescriptors, &out.SpecDescriptors
		*out = make([]SpecDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ActionDescriptor != nil {
		in, out := &in.ActionDescriptor, &out.ActionDescriptor
		*out = make([]ActionDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDDescription.
func (in *CRDDescription) DeepCopy() *CRDDescription {
	if in == nil {
		return nil
	}
	out := new(CRDDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSource) DeepCopyInto(out *CatalogSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSource.
func (in *CatalogSource) DeepCopy() *CatalogSource {
	if in == nil {
		return nil
	}
	out := new(CatalogSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceList) DeepCopyInto(out *CatalogSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceList.
func (in *CatalogSourceList) DeepCopy() *CatalogSourceList {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceSpec) DeepCopyInto(out *CatalogSourceSpec) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Icon = in.Icon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceSpec.
func (in *CatalogSourceSpec) DeepCopy() *CatalogSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceStatus) DeepCopyInto(out *CatalogSourceStatus) {
	*out = *in
	if in.ConfigMapResource != nil {
		in, out := &in.ConfigMapResource, &out.ConfigMapResource
		*out = new(ConfigMapResourceReference)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryServiceStatus != nil {
		in, out := &in.RegistryServiceStatus, &out.RegistryServiceStatus
		*out = new(RegistryServiceStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPCConnectionState != nil {
		in, out := &in.GRPCConnectionState, &out.GRPCConnectionState
		*out = new(GRPCConnectionState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceStatus.
func (in *CatalogSourceStatus) DeepCopy() *CatalogSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceVersion) DeepCopyInto(out *ClusterServiceVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceVersion.
func (in *ClusterServiceVersion) DeepCopy() *ClusterServiceVersion {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceVersionCondition) DeepCopyInto(out *ClusterServiceVersionCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceVersionCondition.
func (in *ClusterServiceVersionCondition) DeepCopy() *ClusterServiceVersionCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceVersionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceVersionList) DeepCopyInto(out *ClusterServiceVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServiceVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceVersionList.
func (in *ClusterServiceVersionList) DeepCopy() *ClusterServiceVersionList {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceVersionSpec) DeepCopyInto(out *ClusterServiceVersionSpec) {
	*out = *in
	in.InstallStrategy.DeepCopyInto(&out.InstallStrategy)
	in.Version.DeepCopyInto(&out.Version)
	in.CustomResourceDefinitions.DeepCopyInto(&out.CustomResourceDefinitions)
	in.APIServiceDefinitions.DeepCopyInto(&out.APIServiceDefinitions)
	if in.NativeAPIs != nil {
		in, out := &in.NativeAPIs, &out.NativeAPIs
		*out = make([]v1.GroupVersionKind, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]Maintainer, len(*in))
		copy(*out, *in)
	}
	out.Provider = in.Provider
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]AppLink, len(*in))
		copy(*out, *in)
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]Icon, len(*in))
		copy(*out, *in)
	}
	if in.InstallModes != nil {
		in, out := &in.InstallModes, &out.InstallModes
		*out = make([]InstallMode, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceVersionSpec.
func (in *ClusterServiceVersionSpec) DeepCopy() *ClusterServiceVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceVersionStatus) DeepCopyInto(out *ClusterServiceVersionStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterServiceVersionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequirementStatus != nil {
		in, out := &in.RequirementStatus, &out.RequirementStatus
		*out = make([]RequirementStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CertsLastUpdated.DeepCopyInto(&out.CertsLastUpdated)
	in.CertsRotateAt.DeepCopyInto(&out.CertsRotateAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceVersionStatus.
func (in *ClusterServiceVersionStatus) DeepCopy() *ClusterServiceVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapResourceReference) DeepCopyInto(out *ConfigMapResourceReference) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapResourceReference.
func (in *ConfigMapResourceReference) DeepCopy() *ConfigMapResourceReference {
	if in == nil {
		return nil
	}
	out := new(ConfigMapResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitions) DeepCopyInto(out *CustomResourceDefinitions) {
	*out = *in
	if in.Owned != nil {
		in, out := &in.Owned, &out.Owned
		*out = make([]CRDDescription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]CRDDescription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitions.
func (in *CustomResourceDefinitions) DeepCopy() *CustomResourceDefinitions {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependentStatus) DeepCopyInto(out *DependentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependentStatus.
func (in *DependentStatus) DeepCopy() *DependentStatus {
	if in == nil {
		return nil
	}
	out := new(DependentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCConnectionState) DeepCopyInto(out *GRPCConnectionState) {
	*out = *in
	in.LastConnectTime.DeepCopyInto(&out.LastConnectTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCConnectionState.
func (in *GRPCConnectionState) DeepCopy() *GRPCConnectionState {
	if in == nil {
		return nil
	}
	out := new(GRPCConnectionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Icon) DeepCopyInto(out *Icon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Icon.
func (in *Icon) DeepCopy() *Icon {
	if in == nil {
		return nil
	}
	out := new(Icon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallMode) DeepCopyInto(out *InstallMode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallMode.
func (in *InstallMode) DeepCopy() *InstallMode {
	if in == nil {
		return nil
	}
	out := new(InstallMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InstallModeSet) DeepCopyInto(out *InstallModeSet) {
	{
		in := &in
		*out = make(InstallModeSet, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallModeSet.
func (in InstallModeSet) DeepCopy() InstallModeSet {
	if in == nil {
		return nil
	}
	out := new(InstallModeSet)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlan) DeepCopyInto(out *InstallPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlan.
func (in *InstallPlan) DeepCopy() *InstallPlan {
	if in == nil {
		return nil
	}
	out := new(InstallPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanCondition) DeepCopyInto(out *InstallPlanCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanCondition.
func (in *InstallPlanCondition) DeepCopy() *InstallPlanCondition {
	if in == nil {
		return nil
	}
	out := new(InstallPlanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanList) DeepCopyInto(out *InstallPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstallPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanList.
func (in *InstallPlanList) DeepCopy() *InstallPlanList {
	if in == nil {
		return nil
	}
	out := new(InstallPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanReference) DeepCopyInto(out *InstallPlanReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanReference.
func (in *InstallPlanReference) DeepCopy() *InstallPlanReference {
	if in == nil {
		return nil
	}
	out := new(InstallPlanReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanSpec) DeepCopyInto(out *InstallPlanSpec) {
	*out = *in
	if in.ClusterServiceVersionNames != nil {
		in, out := &in.ClusterServiceVersionNames, &out.ClusterServiceVersionNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanSpec.
func (in *InstallPlanSpec) DeepCopy() *InstallPlanSpec {
	if in == nil {
		return nil
	}
	out := new(InstallPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanStatus) DeepCopyInto(out *InstallPlanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]InstallPlanCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CatalogSources != nil {
		in, out := &in.CatalogSources, &out.CatalogSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]*Step, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Step)
				**out = **in
			}
		}
	}
	if in.BundleLookups != nil {
		in, out := &in.BundleLookups, &out.BundleLookups
		*out = make([]*BundleLookup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleLookup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AttenuatedServiceAccountRef != nil {
		in, out := &in.AttenuatedServiceAccountRef, &out.AttenuatedServiceAccountRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanStatus.
func (in *InstallPlanStatus) DeepCopy() *InstallPlanStatus {
	if in == nil {
		return nil
	}
	out := new(InstallPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Maintainer) DeepCopyInto(out *Maintainer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Maintainer.
func (in *Maintainer) DeepCopy() *Maintainer {
	if in == nil {
		return nil
	}
	out := new(Maintainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedInstallStrategy) DeepCopyInto(out *NamedInstallStrategy) {
	*out = *in
	if in.StrategySpecRaw != nil {
		in, out := &in.StrategySpecRaw, &out.StrategySpecRaw
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedInstallStrategy.
func (in *NamedInstallStrategy) DeepCopy() *NamedInstallStrategy {
	if in == nil {
		return nil
	}
	out := new(NamedInstallStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryServiceStatus) DeepCopyInto(out *RegistryServiceStatus) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryServiceStatus.
func (in *RegistryServiceStatus) DeepCopy() *RegistryServiceStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequirementStatus) DeepCopyInto(out *RequirementStatus) {
	*out = *in
	if in.Dependents != nil {
		in, out := &in.Dependents, &out.Dependents
		*out = make([]DependentStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequirementStatus.
func (in *RequirementStatus) DeepCopy() *RequirementStatus {
	if in == nil {
		return nil
	}
	out := new(RequirementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecDescriptor) DeepCopyInto(out *SpecDescriptor) {
	*out = *in
	if in.XDescriptors != nil {
		in, out := &in.XDescriptors, &out.XDescriptors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(json.RawMessage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecDescriptor.
func (in *SpecDescriptor) DeepCopy() *SpecDescriptor {
	if in == nil {
		return nil
	}
	out := new(SpecDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusDescriptor) DeepCopyInto(out *StatusDescriptor) {
	*out = *in
	if in.XDescriptors != nil {
		in, out := &in.XDescriptors, &out.XDescriptors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(json.RawMessage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusDescriptor.
func (in *StatusDescriptor) DeepCopy() *StatusDescriptor {
	if in == nil {
		return nil
	}
	out := new(StatusDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	out.Resource = in.Resource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepResource) DeepCopyInto(out *StepResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepResource.
func (in *StepResource) DeepCopy() *StepResource {
	if in == nil {
		return nil
	}
	out := new(StepResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(SubscriptionSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCatalogHealth) DeepCopyInto(out *SubscriptionCatalogHealth) {
	*out = *in
	if in.CatalogSourceRef != nil {
		in, out := &in.CatalogSourceRef, &out.CatalogSourceRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCatalogHealth.
func (in *SubscriptionCatalogHealth) DeepCopy() *SubscriptionCatalogHealth {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCatalogHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionCondition) DeepCopyInto(out *SubscriptionCondition) {
	*out = *in
	if in.LastHeartbeatTime != nil {
		in, out := &in.LastHeartbeatTime, &out.LastHeartbeatTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionCondition.
func (in *SubscriptionCondition) DeepCopy() *SubscriptionCondition {
	if in == nil {
		return nil
	}
	out := new(SubscriptionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionConfig) DeepCopyInto(out *SubscriptionConfig) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionConfig.
func (in *SubscriptionConfig) DeepCopy() *SubscriptionConfig {
	if in == nil {
		return nil
	}
	out := new(SubscriptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		*out = new(InstallPlanReference)
		**out = **in
	}
	if in.InstallPlanRef != nil {
		in, out := &in.InstallPlanRef, &out.InstallPlanRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.CatalogHealth != nil {
		in, out := &in.CatalogHealth, &out.CatalogHealth
		*out = make([]SubscriptionCatalogHealth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]SubscriptionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
