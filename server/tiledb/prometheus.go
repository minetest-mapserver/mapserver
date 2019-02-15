package tiledb

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	setDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiledb_set_duration",
		Help:    "Histogram for tiledb set timings",
		Buckets: prometheus.LinearBuckets(0.001, 0.01, 10),
	})
	getDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiledb_get_duration",
		Help:    "Histogram for tiledb get timings",
		Buckets: prometheus.LinearBuckets(0.001, 0.01, 10),
	})
	removeDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiledb_remove_duration",
		Help:    "Histogram for tiledb remove timings",
		Buckets: prometheus.LinearBuckets(0.001, 0.01, 10),
	})
)

func init() {
	prometheus.MustRegister(setDuration)
	prometheus.MustRegister(getDuration)
	prometheus.MustRegister(removeDuration)
}
