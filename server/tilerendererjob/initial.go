package tilerendererjob

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"time"
)

func initialRender(ctx *app.App, jobs chan *coords.TileCoords) {

	rstate := ctx.Config.RenderState
	totalLegacyCount, err := ctx.Blockdb.CountBlocks(0, 0)

	if err != nil {
		panic(err)
	}

	fields := logrus.Fields{
		"totalLegacyCount": totalLegacyCount,
		"LastMtime":        rstate.LastMtime,
	}
	logrus.WithFields(fields).Info("Starting initial rendering job")

	lastcoords := coords.NewMapBlockCoords(rstate.LastX, rstate.LastY, rstate.LastZ)

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindMapBlocksByPos(lastcoords, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			rstate.InitialRun = false
			ctx.Config.Save()

			fields := logrus.Fields{
				"legacyblocks": rstate.LegacyProcessed,
			}
			logrus.WithFields(fields).Info("initial rendering complete")

			return
		}

		renderMapblocks(ctx, jobs, result.List)

		lastcoords = *result.LastPos
		rstate.LastMtime = result.LastMtime

		//Save current positions of initial run
		rstate.LastX = lastcoords.X
		rstate.LastY = lastcoords.Y
		rstate.LastZ = lastcoords.Z
		rstate.LegacyProcessed += result.UnfilteredCount
		ctx.Config.Save()

		t := time.Now()
		elapsed := t.Sub(start)

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

	}
}
