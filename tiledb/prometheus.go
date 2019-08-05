package tiledb

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	tiledbSaveDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiledb_save_durations",
		Help:    "Histogram for tiledb save timings",
		Buckets: prometheus.LinearBuckets(0.005, 0.01, 10),
	})
	tiledbLoadDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiledb_load_durations",
		Help:    "Histogram for tiledb load timings",
		Buckets: prometheus.LinearBuckets(0.005, 0.01, 10),
	})
)

func init() {
	prometheus.MustRegister(tiledbSaveDuration)
	prometheus.MustRegister(tiledbLoadDuration)
}
