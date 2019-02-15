package mapblockrenderer

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	renderedMapblocks = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "mapblock_rendered_count",
			Help: "Overall count of rendered mapblocks",
		},
	)
	renderDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "mapblock_render_time",
		Help:    "Histogram for mapblock render timings",
		Buckets: prometheus.LinearBuckets(0.01, 0.02, 10),
	})
)

func init() {
	prometheus.MustRegister(renderedMapblocks)
	prometheus.MustRegister(renderDuration)
}
