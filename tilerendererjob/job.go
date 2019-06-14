package tilerendererjob

import (
	"mapserver/app"
	"mapserver/settings"
)

func Job(ctx *app.App) {
	lastMtime := ctx.Settings.GetInt64(settings.SETTING_LAST_MTIME, 0)
	if lastMtime == 0 {
		//mark db time as last incremental render point
		lastMtime, err := ctx.Blockdb.GetTimestamp()

		if err != nil {
			panic(err)
		}

		ctx.Settings.SetInt64(settings.SETTING_LAST_MTIME, lastMtime)
	}

	if ctx.Config.EnableInitialRendering {
		if ctx.Settings.GetBool(settings.SETTING_INITIAL_RUN, true) {
			initialRender(ctx)
		}
	}

	incrementalRender(ctx)

	panic("render job interrupted!")

}
