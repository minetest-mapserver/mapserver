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

var blankTile = tilerenderer.CreateBlankTile(color.RGBA{255, 255, 255, 255})

type Tiles struct {
	ctx *app.App
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
		resp.Header().Add("Content-Type", "image/png")

		if tile == nil {
			// cache blank tile for a while (heavy re-use)
			resp.Header().Add("Cache-Control", "max-age=300")
			resp.Write(blankTile)

		} else {
			// cache tile up to 10 seconds (realtime)
			resp.Header().Add("Cache-Control", "max-age=10")
			resp.Write(tile)

		}
	}
}
