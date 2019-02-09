package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/settings"
)

func Job(ctx *app.App) {
	initMetrics()

	jobs := make(chan *coords.TileCoords, ctx.Config.RenderingQueue)

	for i := 0; i < ctx.Config.RenderingJobs; i++ {
		go worker(ctx, jobs)
	}

	if ctx.Settings.GetBool(settings.SETTING_INITIAL_RUN, true) {
		initialRender(ctx, jobs)
	}

	incrementalRender(ctx, jobs)

	panic("render job interrupted!")

}
