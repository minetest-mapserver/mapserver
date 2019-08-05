package web

import (
	"github.com/prometheus/client_golang/prometheus"
	"image/color"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/tilerenderer"
	"net/http"
	"strconv"
	"strings"
)

type Tiles struct {
	ctx   *app.App
	blank []byte
}

func (t *Tiles) Init() {
	t.blank = tilerenderer.CreateBlankTile(color.RGBA{255, 255, 255, 255})
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

	timer := prometheus.NewTimer(tileServeDuration)
	defer timer.ObserveDuration()

	layerid, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	zoom, _ := strconv.Atoi(parts[3])

	c := coords.NewTileCoords(x, y, zoom, layerid)
	tile, err := t.ctx.TileDB.GetTile(c)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {
		resp.Header().Add("content-type", "image/png")

		if tile == nil {
			resp.Write(t.blank)
			//TODO: cache/layer color

		} else {
			tilesCumulativeSize.Add(float64(len(tile)))
			resp.Write(tile)

		}
	}
}
