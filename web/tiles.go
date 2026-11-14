package web

import (
	"image/color"
	"mapserver/app"
	"mapserver/tilerenderer"
	"mapserver/util"
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
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("wrong number of arguments"))
		return
	}

	timer := prometheus.NewTimer(tileServeDuration)
	defer timer.ObserveDuration()

	layerid, err1 := strconv.Atoi(parts[0])
	x, err2 := strconv.Atoi(parts[1])
	y, err3 := strconv.Atoi(parts[2])
	zoom, err4 := strconv.Atoi(parts[3])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid tile coordinates"))
		return
	}

	c := util.NewTileCoords(x, y, zoom, layerid)
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
