package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
)

func Job(ctx *app.App) {

	rstate := ctx.Config.RenderState
	jobs := make(chan *coords.TileCoords, ctx.Config.RenderingQueue)

	for i := 0; i < ctx.Config.RenderingJobs; i++ {
		go worker(ctx, jobs)
	}

	if rstate.InitialRun {
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
