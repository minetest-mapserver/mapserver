package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/settings"
	"time"

	"github.com/sirupsen/logrus"
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
		"totalLegacyCount": totalLegacyCount,
	}
	logrus.WithFields(fields).Info("Starting initial rendering job")

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindNextLegacyBlocks(ctx.Settings, ctx.Config.Layers, ctx.Config.RenderingFetchLimit)

		if err != nil {
			panic(err)
		}

		legacyProcessed := ctx.Settings.GetInt(settings.SETTING_LEGACY_PROCESSED, 0)

		if len(result.List) == 0 && !result.HasMore {
			ctx.Settings.SetBool(settings.SETTING_INITIAL_RUN, false)

			ev := InitialRenderEvent{
				Progress: 100,
			}

			ctx.WebEventbus.Emit("initial-render-progress", &ev)

			fields := logrus.Fields{
				"legacyblocks": legacyProcessed,
			}
			logrus.WithFields(fields).Info("initial rendering complete")

			return
		}

		tiles := renderMapblocks(ctx, jobs, result.List)

		legacyProcessed += result.UnfilteredCount

		ctx.Settings.SetInt(settings.SETTING_LEGACY_PROCESSED, legacyProcessed)

		t := time.Now()
		elapsed := t.Sub(start)

		progress := int(float64(legacyProcessed) / float64(totalLegacyCount) * 100)

		ev := InitialRenderEvent{
			Progress: progress,
		}

		ctx.WebEventbus.Emit("initial-render-progress", &ev)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
			"processed": legacyProcessed,
			"progress%": progress,
			"elapsed":   elapsed,
		}
		logrus.WithFields(fields).Info("Initial rendering")

		//tile gc
		ctx.TileDB.GC()

	}
}
