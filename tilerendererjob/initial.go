package tilerendererjob

import (
	"mapserver/app"
	"mapserver/settings"
	"time"

	"github.com/sirupsen/logrus"
)

type InitialRenderEvent struct {
	Progress float64 `json:"progress"`
}

func initialRender(ctx *app.App) {
	logrus.Info("Starting initial rendering job")

	for true {
		start := time.Now()

		result, err := ctx.MapBlockAccessor.FindNextLegacyBlocks(ctx.Settings, ctx.Config.Layers, ctx.Config.RenderingFetchLimit)

		if err != nil {
			logrus.Error("Error in initial rendering run, trying to continue: " + err.Error())
			continue
		}

		if len(result.List) == 0 && !result.HasMore {
			ctx.Settings.SetBool(settings.SETTING_INITIAL_RUN, false)

			ev := InitialRenderEvent{
				Progress: 1,
			}

			ctx.WebEventbus.Emit("initial-render-progress", &ev)

			logrus.Info("initial rendering complete")

			return
		}

		tiles := renderMapblocks(ctx, result.List)

		t := time.Now()
		elapsed := t.Sub(start)

		ev := InitialRenderEvent{
			Progress: result.Progress,
		}

		ctx.WebEventbus.Emit("initial-render-progress", &ev)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
			"progress%": int(result.Progress * 100),
			"elapsed":   elapsed,
		}
		logrus.WithFields(fields).Info("Initial rendering")

		//tile gc
		ctx.TileDB.GC()

	}
}
