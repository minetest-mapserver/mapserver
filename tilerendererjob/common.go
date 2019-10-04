package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapblockparser"
	"strconv"
	"github.com/sirupsen/logrus"
)

const(
	MAX_UNZOOM = 13
)

func getTileKey(tc *coords.TileCoords) string {
	return strconv.Itoa(tc.X) + "/" + strconv.Itoa(tc.Y) + "/" +
		strconv.Itoa(tc.Zoom) + "/" + strconv.Itoa(tc.LayerId)
}

func renderMapblocks(ctx *app.App, mblist []*mapblockparser.MapBlock) int {
	tilecount := 0
	totalRenderedMapblocks.Add(float64(len(mblist)))

	for _, mb := range mblist {
		tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)
		ctx.TileDB.MarkOutdated(tc)
	}

	for uz := 0; uz <= MAX_UNZOOM; uz++ {
		//Spin up workers
		jobs := make(chan coords.TileCoords, ctx.Config.RenderingQueue)
		done := make(chan bool, 1)

		for j := 0; j < ctx.Config.RenderingJobs; j++ {
			go worker(ctx, jobs, done)
		}

		for _, tc := range ctx.TileDB.GetOutdatedByZoom(uz) {
			fields := logrus.Fields{
				"pos":    tc,
				"prefix": "tilerenderjob",
			}
			logrus.WithFields(fields).Debug("Tile render job mapblock")

			tilecount++

			fields = logrus.Fields{
				"X":       tc.X,
				"Y":       tc.Y,
				"Zoom":    tc.Zoom,
				"LayerId": tc.LayerId,
				"prefix":  "tilerenderjob",
			}
			logrus.WithFields(fields).Debug("Tile render job dispatch tile")

			//dispatch re-render
			jobs <- tc
		}
		//spin down worker pool
		close(jobs)

		for j := 0; j < ctx.Config.RenderingJobs; j++ {
			<-done
		}
	}

	return tilecount
}
