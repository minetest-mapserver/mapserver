package tilerenderer

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	renderedTiles = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tiles_rendered_count",
			Help: "Overall count of rendered tiles",
		},
		[]string{"zoom"},
	)
	renderDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiles_render_time",
		Help:    "Histogram for tile render timings",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	})
)

func init() {
	prometheus.MustRegister(renderedTiles)
	prometheus.MustRegister(renderDuration)
}
