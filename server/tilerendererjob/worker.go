package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
)

func worker(ctx *app.App, coords <-chan *coords.TileCoords) {
	for tc := range coords {
		ctx.Objectdb.RemoveTile(tc)
		_, err := ctx.Tilerenderer.Render(tc, 2)
		if err != nil {
			panic(err)
		}
	}
}
