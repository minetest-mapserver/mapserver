package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
)

func worker(ctx *app.App, coords <-chan *coords.TileCoords) {
	for tc := range coords {
		_, err := ctx.Tilerenderer.Render(tc, 5)
		if err != nil {
			panic(err)
		}
	}
}
