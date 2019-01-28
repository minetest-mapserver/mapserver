package mapobject

import (
	"mapserver/app"
	"mapserver/eventbus"
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type MapObjectListener interface {
	onMapObject(x,y,z int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor)
}

type Listener struct {
	ctx *app.App
	objectlisteners map[string]MapObjectListener
}

func (this *Listener) AddMapObject(blockname string, ol MapObjectListener){
	this.objectlisteners[blockname] = ol
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
		for k, v := range this.objectlisteners {
			if k == name {
				//block matches

				for x := 0; x < 16; x++ {
					for y := 0; y < 16; y++ {
						for z := 0; z < 16; z++ {
							nodeid := block.GetNodeId(x, y, z)
							if nodeid == id {
								v.onMapObject(x, y, z, block, this.ctx.Objectdb)
							}
						}//z
					}//y
				}//x

			}
		}//for k,v
	}//for id, name
}

func Setup(ctx *app.App) {
	l := Listener{
		ctx: ctx,
		objectlisteners: make(map[string]MapObjectListener),
	}

	l.AddMapObject("mapserver:poi", &PoiBlock{})
	l.AddMapObject("mapserver:train", &TrainBlock{})
	l.AddMapObject("travelnet:travelnet", &TravelnetBlock{})
	l.AddMapObject("protector:protect", &ProtectorBlock{})
	l.AddMapObject("protector:protect2", &ProtectorBlock{})

	ctx.BlockAccessor.Eventbus.AddListener(&l)
}
