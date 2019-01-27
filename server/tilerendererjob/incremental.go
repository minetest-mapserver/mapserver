package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"time"

	"github.com/sirupsen/logrus"
)

func incrementalRender(ctx *app.App, jobs chan *coords.TileCoords) {

	rstate := ctx.Config.RenderState

	fields := logrus.Fields{
		"LastMtime": rstate.LastMtime,
	}
	logrus.WithFields(fields).Info("Starting incremental rendering job")

	for true {
		start := time.Now()

		result, err := ctx.BlockAccessor.FindMapBlocksByMtime(rstate.LastMtime, ctx.Config.RenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(result.List) == 0 && !result.HasMore {
			time.Sleep(5 * time.Second)
			continue
		}

		rstate.LastMtime = result.LastMtime
		ctx.Config.Save()

		tiles := renderMapblocks(ctx, jobs, result.List)

		t := time.Now()
		elapsed := t.Sub(start)

		fields := logrus.Fields{
			"mapblocks": len(result.List),
			"tiles":     tiles,
			"elapsed":   elapsed,
		}
		logrus.WithFields(fields).Info("incremental rendering")
	}
}
