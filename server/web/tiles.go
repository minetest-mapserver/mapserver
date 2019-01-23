package web

import (
	"image/color"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/tilerenderer"
	"net/http"
	"strconv"
	"strings"
)

type Tiles struct {
	ctx *app.App
}

func (t *Tiles) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	str := strings.TrimPrefix(req.URL.Path, "/api/tile/")
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
			resp.Write(tilerenderer.CreateBlankTile(color.RGBA{0, 0, 0, 0}))
			//TODO: cache/layer color

		} else {
			resp.Write(tile.Data)

		}
	}
}
