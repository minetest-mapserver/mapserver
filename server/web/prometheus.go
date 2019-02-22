package web

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	tilesCumulativeSize = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tiles_cumulative_size_served",
			Help: "Overall sent bytes of tiles",
		},
	)
	tileServeDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tiles_serve_durations",
		Help:    "Histogram for tile serve timings",
		Buckets: prometheus.LinearBuckets(0.005, 0.01, 10),
	})
	mapobjectServeDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "mapobject_serve_durations",
		Help:    "Histogram for mapobject serve timings",
		Buckets: prometheus.LinearBuckets(0.005, 0.01, 10),
	})
	wsClients = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ws_client_count",
		Help: "Websocket client count",
	})
	mintestPlayers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mapserver_minetest_player_count",
		Help: "game player count",
	})
	mintestMaxLag = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mapserver_minetest_max_lag",
		Help: "Max lag",
	})
)

func init() {
	prometheus.MustRegister(tilesCumulativeSize)
	prometheus.MustRegister(tileServeDuration)
	prometheus.MustRegister(mapobjectServeDuration)
	prometheus.MustRegister(wsClients)
	prometheus.MustRegister(mintestPlayers)
	prometheus.MustRegister(mintestMaxLag)
}
