package tilerendererjob

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"time"
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

		if len(result.List) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		renderMapblocks(ctx, jobs, result.List)

		rstate.LastMtime = result.LastMtime
		ctx.Config.Save()

		t := time.Now()
		elapsed := t.Sub(start)

		fields := logrus.Fields{
			"count":   len(result.List),
			"elapsed": elapsed,
		}
		logrus.WithFields(fields).Info("incremental rendering")
	}
}
