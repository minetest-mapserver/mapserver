package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
)

func worker(ctx *app.App, coords <-chan *coords.TileCoords) {
	for tc := range coords {
		//render tile
		_, err := ctx.Tilerenderer.Render(tc)
		if err != nil {
			panic(err)
		}
	}
}
