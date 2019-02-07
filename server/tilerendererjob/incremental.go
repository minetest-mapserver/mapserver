package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/settings"
	"time"

	"github.com/sirupsen/logrus"
)

type IncrementalRenderEvent struct {
	LastMtime int64 `json:"lastmtime"`
}

func incrementalRender(ctx *app.App, jobs chan *coords.TileCoords) {

	lastMtime := ctx.Settings.GetInt64(settings.SETTING_LAST_MTIME, 0)

	fields := logrus.Fields{
		"LastMtime": lastMtime,
	}
	logrus.WithFields(fields).Info("Starting incremental rendering job")

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindMapBlocksByMtime(lastMtime, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			time.Sleep(5 * time.Second)
			continue
		}

		lastMtime = result.LastMtime
		ctx.Settings.SetInt64(settings.SETTING_LAST_MTIME, lastMtime)

		tiles := renderMapblocks(ctx, jobs, result.List)

		t := time.Now()
		elapsed := t.Sub(start)

		ev := IncrementalRenderEvent{
			LastMtime: result.LastMtime,
		}

		ctx.WebEventbus.Emit("incremental-render-progress", &ev)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
			"elapsed":   elapsed,
		}
		logrus.WithFields(fields).Info("incremental rendering")
	}
}
