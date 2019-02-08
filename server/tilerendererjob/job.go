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
		//fast, unsafe mode
		err := ctx.Objectdb.EnableSpeedSafetyTradeoff(true)
		if err != nil {
			panic(err)
		}

		initialRender(ctx, jobs)

		//normal, safe mode
		err = ctx.Objectdb.EnableSpeedSafetyTradeoff(false)
		if err != nil {
			panic(err)
		}

	}

	incrementalRender(ctx, jobs)

	panic("render job interrupted!")

}
