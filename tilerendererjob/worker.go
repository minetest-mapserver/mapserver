package tilerendererjob

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
)

func worker(ctx *app.App, coords <-chan *coords.TileCoords, done chan bool) {
	for tc := range coords {
		//render tile

		fields := logrus.Fields{
			"X":       tc.X,
			"Y":       tc.Y,
			"Zoom":    tc.Zoom,
			"LayerId": tc.LayerId,
			"prefix":  "tilerenderjob",
		}
		logrus.WithFields(fields).Debug("Tile render job tile")

		err := ctx.Tilerenderer.Render(tc)
		if err != nil {
			fields := logrus.Fields{
				"X":       tc.X,
				"Y":       tc.Y,
				"Zoom":    tc.Zoom,
				"LayerId": tc.LayerId,
				"prefix":  "tilerenderjob",
				"err":     err,
			}
			logrus.WithFields(fields).Error("Tile render job tile")
		}
	}

	done <- true
}
