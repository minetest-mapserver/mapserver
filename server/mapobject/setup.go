package mapobject

import (
	"mapserver/app"
	"mapserver/eventbus"
	"mapserver/mapblockparser"
)

type Listener struct {
	ctx *app.App
}

func (this *Listener) OnEvent(eventtype string, o interface{}) {
	if eventtype != eventbus.MAPBLOCK_RENDERED {
		return
	}

	block := o.(*mapblockparser.MapBlock)

	err := this.ctx.Objectdb.RemoveMapData(&block.Pos)
	if err != nil {
		panic(err)
	}

	for id, name := range block.BlockMapping {
		if name == "mapserver:poi" {
			onPoiBlock(id, block, this.ctx.Objectdb)
		}
	}
}

func Setup(ctx *app.App) {
	ctx.BlockAccessor.Eventbus.AddListener(&Listener{ctx: ctx})
}
