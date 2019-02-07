package tilerendererjob

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/settings"
	"mapserver/coords"
	"time"
)

type InitialRenderEvent struct {
	Progress int `json:"progress"`
}

func initialRender(ctx *app.App, jobs chan *coords.TileCoords) {

	totalLegacyCount, err := ctx.Blockdb.CountBlocks(0, 0)

	if err != nil {
		panic(err)
	}

	fields := logrus.Fields{
		"totalLegacyCount": totalLegacyCount
	}
	logrus.WithFields(fields).Info("Starting initial rendering job")

	lastx := ctx.Settings.GetInt(settings.SETTING_LASTX, -1)
	lasty := ctx.Settings.GetInt(settings.SETTING_LASTY, -1)
	lastz := ctx.Settings.GetInt(settings.SETTING_LASTZ, -1)

	lastcoords := coords.NewMapBlockCoords(lastx, lasty, lastz)

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindMapBlocksByPos(lastcoords, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			ctx.Settings.SetBool(settings.SETTING_INITIAL_RUN, false)

			ev := InitialRenderEvent{
				Progress: 100,
			}

			ctx.WebEventbus.Emit("initial-render-progress", &ev)

			fields := logrus.Fields{
				"legacyblocks": rstate.LegacyProcessed,
			}
			logrus.WithFields(fields).Info("initial rendering complete")

			return
		}

		tiles := renderMapblocks(ctx, jobs, result.List)

		lastcoords = result.LastPos
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

		ev := InitialRenderEvent{
			Progress: progress,
		}

		ctx.WebEventbus.Emit("initial-render-progress", &ev)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
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
