package tilerendererjob

import (
	"mapserver/app"
	"mapserver/settings"
)

func Job(ctx *app.App) {
	if ctx.Settings.GetBool(settings.SETTING_INITIAL_RUN, true) {
		initialRender(ctx)
	}

	incrementalRender(ctx)

	panic("render job interrupted!")

}
