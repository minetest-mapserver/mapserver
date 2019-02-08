package web

import (
	"image/color"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/tilerenderer"
	"net/http"
	"strconv"
	"strings"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	tilesCumulativeSize = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tiles_cumulative_size_served",
			Help: "Overall sent bytes of tiles",
		},
	)
)

type Tiles struct {
	ctx   *app.App
	blank []byte
}

func (t *Tiles) Init() {
	t.blank = tilerenderer.CreateBlankTile(color.RGBA{255, 255, 255, 255})
	prometheus.MustRegister(tilesCumulativeSize)
}

func (t *Tiles) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/tile/")
	// {layerId}/x/y/zoom
	parts := strings.Split(str, "/")
	if len(parts) != 4 {
		resp.WriteHeader(500)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	layerid, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	zoom, _ := strconv.Atoi(parts[3])

	c := coords.NewTileCoords(x, y, zoom, layerid)
	tile, err := t.ctx.Objectdb.GetTile(c)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {
		resp.Header().Add("content-type", "image/png")

		if tile == nil {
			resp.Write(t.blank)
			//TODO: cache/layer color

		} else {
			tilesCumulativeSize.Add(float64(len(tile.Data)))
			resp.Write(tile.Data)

		}
	}
}
