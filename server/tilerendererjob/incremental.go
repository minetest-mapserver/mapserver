package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapobjectdb"
	"time"
	"strconv"

	"github.com/sirupsen/logrus"
)

type IncrementalRenderEvent struct {
	LastMtime int64 `json:"lastmtime"`
}

func incrementalRender(ctx *app.App, jobs chan *coords.TileCoords) {

	str, err := ctx.Objectdb.GetSetting(mapobjectdb.SETTING_LAST_MTIME, "0")
	if err != nil {
		panic(err)
	}

	lastMtime, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

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
		ctx.Objectdb.SetSetting(mapobjectdb.SETTING_LAST_MTIME, strconv.FormatInt(lastMtime, 10))

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
