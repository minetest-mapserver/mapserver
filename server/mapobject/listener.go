package mapobject

import (
	"mapserver/app"
	"mapserver/eventbus"
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"

	"github.com/sirupsen/logrus"
)

type MapObjectListener interface {
	onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject
}

type MapMultiObjectListener interface {
	onMapObject(x, y, z int, block *mapblockparser.MapBlock) []*mapobjectdb.MapObject
}

type Listener struct {
	ctx                  *app.App
	objectlisteners      map[string]MapObjectListener
	multiobjectlisteners map[string]MapMultiObjectListener
}

func (this *Listener) AddMapObject(blockname string, ol MapObjectListener) {
	this.objectlisteners[blockname] = ol
}

func (this *Listener) AddMapMultiObject(blockname string, ol MapMultiObjectListener) {
	this.multiobjectlisteners[blockname] = ol
}

func (this *Listener) OnEvent(eventtype string, o interface{}) {
	if eventtype != eventbus.MAPBLOCK_RENDERED {
		return
	}

	block := o.(*mapblockparser.MapBlock)

	err := this.ctx.Objectdb.RemoveMapData(block.Pos)
	if err != nil {
		panic(err)
	}

	this.ctx.WebEventbus.Emit("mapobjects-cleared", block.Pos)

	for id, name := range block.BlockMapping {

		for k, v := range this.multiobjectlisteners {
			if k == name {
				//block matches
				mapblockparser.IterateMapblock(func(x, y, z int) {
					nodeid := block.GetNodeId(x, y, z)
					if nodeid == id {
						fields := logrus.Fields{
							"mbpos":  block.Pos,
							"x":      x,
							"y":      y,
							"z":      z,
							"type":   name,
							"nodeid": nodeid,
						}
						log.WithFields(fields).Debug("OnEvent()")

						objs := v.onMapObject(x, y, z, block)

						if len(objs) > 0 {
							for _, obj := range objs {
								err := this.ctx.Objectdb.AddMapData(obj)
								if err != nil {
									panic(err)
								}

								this.ctx.WebEventbus.Emit("mapobject-created", obj)
							}
						}
					}
				})
			} // k==name
		} //for k,v

		for k, v := range this.objectlisteners {
			if k == name {
				//block matches
				mapblockparser.IterateMapblock(func(x, y, z int) {
					nodeid := block.GetNodeId(x, y, z)
					if nodeid == id {
						fields := logrus.Fields{
							"mbpos":  block.Pos,
							"x":      x,
							"y":      y,
							"z":      z,
							"type":   name,
							"nodeid": nodeid,
						}
						log.WithFields(fields).Debug("OnEvent()")

						obj := v.onMapObject(x, y, z, block)

						if obj != nil {
							this.ctx.Objectdb.AddMapData(obj)
							this.ctx.WebEventbus.Emit("mapobject-created", obj)
						}
					}
				})
			} // k==name
		} //for k,v

	} //for id, name
}
