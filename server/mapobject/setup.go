package mapobject

import (
	"mapserver/app"
)

func Setup(ctx *app.App) {
	ctx.BlockAccessor.AddListener(&ClearMapData{db: ctx.Objectdb})
	ctx.BlockAccessor.AddListener(&POI{db: ctx.Objectdb})
}
