package tilerendererjob

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"strconv"
	"time"
)

func getTileKey(tc *coords.TileCoords) string {
	return strconv.Itoa(tc.X) + "/" + strconv.Itoa(tc.Y) + "/" + strconv.Itoa(tc.Zoom)
}

func worker(ctx *app.App, coords <-chan *coords.TileCoords) {
	for tc := range coords {
		ctx.Objectdb.RemoveTile(tc)
		_, err := ctx.Tilerenderer.Render(tc, 2)
		if err != nil {
			panic(err)
		}
	}
}

func Job(ctx *app.App) {

	rstate := ctx.Config.RenderState
	var totalLegacyCount int
	var err error

	if rstate.InitialRun {
		totalLegacyCount, err = ctx.Blockdb.CountBlocks(0, 0)

		if err != nil {
			panic(err)
		}

		fields := logrus.Fields{
			"totalLegacyCount": totalLegacyCount,
			"LastMtime":        rstate.LastMtime,
		}
		logrus.WithFields(fields).Info("Starting rendering job")

	} else {
		fields := logrus.Fields{
			"LastMtime": rstate.LastMtime,
		}
		logrus.WithFields(fields).Info("Starting rendering job")

	}

	tilecount := 0

	lastcoords := coords.NewMapBlockCoords(rstate.LastX, rstate.LastY, rstate.LastZ)

	jobs := make(chan *coords.TileCoords, ctx.Config.RenderingQueue)

	for i := 0; i < ctx.Config.RenderingJobs; i++ {
		go worker(ctx, jobs)
	}

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindMapBlocks(lastcoords, rstate.LastMtime, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			if rstate.InitialRun {
				rstate.InitialRun = false
				ctx.Config.Save()

				fields := logrus.Fields{
					"legacyblocks": rstate.LegacyProcessed,
				}
				logrus.WithFields(fields).Info("initial rendering complete")
			}

			time.Sleep(5 * time.Second)
			continue
		}

		tileRenderedMap := make(map[string]bool)

		for i := 12; i >= 1; i-- {
			for _, mb := range result.List {
				//13
				tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

				//12-1
				tc = tc.ZoomOut(13 - i)

				key := getTileKey(tc)

				if tileRenderedMap[key] {
					continue
				}

				tileRenderedMap[key] = true

				fields := logrus.Fields{
					"X":       tc.X,
					"Y":       tc.Y,
					"Zoom":    tc.Zoom,
					"LayerId": tc.LayerId,
				}
				logrus.WithFields(fields).Debug("Dispatching tile rendering (z11-1)")

				tilecount++
				jobs <- tc
			}
		}

		lastcoords = *result.LastPos
		rstate.LastMtime = result.LastMtime

		if rstate.InitialRun {
			//Save current positions of initial run
			rstate.LastX = lastcoords.X
			rstate.LastY = lastcoords.Y
			rstate.LastZ = lastcoords.Z
			rstate.LegacyProcessed += result.UnfilteredCount
		}

		ctx.Config.Save()

		t := time.Now()
		elapsed := t.Sub(start)

		if rstate.InitialRun {
			progress := int(float64(rstate.LegacyProcessed) / float64(totalLegacyCount) * 100)

			fields := logrus.Fields{
				"count":     len(result.List),
				"processed": rstate.LegacyProcessed,
				"progress%": progress,
				"X":         lastcoords.X,
				"Y":         lastcoords.Y,
				"Z":         lastcoords.Z,
				"elapsed":   elapsed,
			}
			logrus.WithFields(fields).Info("Initial rendering")

		} else {
			fields := logrus.Fields{
				"count":   len(result.List),
				"elapsed": elapsed,
			}
			logrus.WithFields(fields).Info("incremental rendering")

		}
	}
}
