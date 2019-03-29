package tilerendererjob

import (
	"mapserver/app"
	"mapserver/settings"
	"time"
)

func Job(ctx *app.App) {
	lastMtime := ctx.Settings.GetInt64(settings.SETTING_LAST_MTIME, 0)
	if lastMtime == 0 {
		//mark current time as last incremental render point
		ctx.Settings.SetInt64(settings.SETTING_LAST_MTIME, time.Now().Unix())
	}

	if ctx.Config.EnableInitialRendering {
		if ctx.Settings.GetBool(settings.SETTING_INITIAL_RUN, true) {
			initialRender(ctx)
		}
	}

	incrementalRender(ctx)

	panic("render job interrupted!")

}
