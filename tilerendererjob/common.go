package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"strconv"

	"github.com/minetest-go/mapparser"
	"github.com/sirupsen/logrus"
)

func getTileKey(tc *coords.TileCoords) string {
	return strconv.Itoa(tc.X) + "/" + strconv.Itoa(tc.Y) + "/" +
		strconv.Itoa(tc.Zoom) + "/" + strconv.Itoa(tc.LayerId)
}

func renderMapblocks(ctx *app.App, mblist []*mapparser.MapBlock) int {
	tileRenderedMap := make(map[string]bool)
	tilecount := 0
	totalRenderedMapblocks.Add(float64(len(mblist)))

	for i := 12; i >= 1; i-- {

		//Spin up workers
		jobs := make(chan *coords.TileCoords, ctx.Config.RenderingQueue)
		done := make(chan bool, 1)

		for j := 0; j < ctx.Config.RenderingJobs; j++ {
			go worker(ctx, jobs, done)
		}

		for _, mb := range mblist {
			//13

			fields := logrus.Fields{
				"pos":    mb.Pos,
				"prefix": "tilerenderjob",
			}
			logrus.WithFields(fields).Debug("Tile render job mapblock")

			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			if tc == nil {
				fields := logrus.Fields{
					"pos":  mb.Pos,
					"zoom": i,
				}
				logrus.WithFields(fields).Error("mapblock outside of layer!")
				panic("mapblock outside of layer!")
			}

			//12-1
			tc = tc.ZoomOut(13 - i)

			key := getTileKey(tc)

			if tileRenderedMap[key] {
				continue
			}

			tileRenderedMap[key] = true

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
