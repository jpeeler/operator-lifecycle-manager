package metrics

import "github.com/prometheus/client_golang/prometheus"

// To add new metrics:
// 1. register new metrics in main.go<olm>
// 2. add appropriate metric updates in handleMetrics
var (
	CSVCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "csv_count",
			Help: "Number of CSVs successfully registered",
		},
	)

	InstallPlanCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "install_plan_count",
			Help: "Number of install plans",
		},
	)

	SubscriptionCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "subscription_count",
			Help: "Number of subscriptions",
		},
	)

	CatalogSourceCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "catalog_source_count",
			Help: "Number of catalog sources",
		},
	)

	CSVUpgradeCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "csv_upgrade_count",
			Help: "Monotonic count of catalog sources",
		},
	)
)
