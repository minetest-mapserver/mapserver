package mapblockparser

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	parsedMapBlocks = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "mapblocks_parsed_count",
			Help: "Overall count of parsed mapblocks",
		},
	)
	parseDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "mapblock_parse_time",
		Help:    "Histogram for mapblock parse timings",
		Buckets: prometheus.LinearBuckets(0.001, 0.002, 10),
	})
)

func init() {
	prometheus.MustRegister(parsedMapBlocks)
	prometheus.MustRegister(parseDuration)
}
