package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"github.com/sirupsen/logrus"
)

func worker(ctx *app.App, coords <-chan *coords.TileCoords) {
	for tc := range coords {
		//render tile

		fields := logrus.Fields{
			"X":       tc.X,
			"Y":       tc.Y,
			"Zoom":    tc.Zoom,
			"LayerId": tc.LayerId,
			"prefix": "tilerenderjob",
		}
		logrus.WithFields(fields).Debug("Tile render job tile")

		_, err := ctx.Tilerenderer.Render(tc)
		if err != nil {
			panic(err)
		}
	}
}
