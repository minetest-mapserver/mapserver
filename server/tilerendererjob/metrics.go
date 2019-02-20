package tilerendererjob

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRenderedMapblocks = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_rendered_mapblocks",
			Help: "Overall rendered mapblocks",
		},
	)
)

func init() {
	prometheus.MustRegister(totalRenderedMapblocks)
}
