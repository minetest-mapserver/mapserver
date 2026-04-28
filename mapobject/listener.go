package mapobject

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/eventbus"
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
	"github.com/sirupsen/logrus"
)

type MapObjectListener interface {
	onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject
}

type MapMultiObjectListener interface {
	onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) []*mapobjectdb.MapObject
}

type Listener struct {
	ctx                  *app.App
	objectlisteners      map[string]MapObjectListener
	multiobjectlisteners map[string]MapMultiObjectListener
}

func (l *Listener) AddMapObject(blockname string, ol MapObjectListener) {
	l.objectlisteners[blockname] = ol
}

func (l *Listener) AddMapMultiObject(blockname string, ol MapMultiObjectListener) {
	l.multiobjectlisteners[blockname] = ol
}

func (l *Listener) OnEvent(eventtype string, o interface{}) {
	if eventtype != eventbus.MAPBLOCK_RENDERED {
		return
	}

	pmb := o.(*types.ParsedMapblock)

	err := l.ctx.Objectdb.RemoveMapData(pmb.Pos)
	if err != nil {
		panic(err)
	}

	l.ctx.WebEventbus.Emit("mapobjects-cleared", pmb.Pos)

	//TODO: refactor into single loop
	for id, name := range pmb.Mapblock.BlockMapping {

		for k, v := range l.multiobjectlisteners {
			if k == name {
				//block matches
				coords.IterateMapblock(func(x, y, z int) {
					nodeid := pmb.Mapblock.GetNodeId(x, y, z)
					if nodeid == id {
						fields := logrus.Fields{
							"mbpos":  pmb.Pos,
							"x":      x,
							"y":      y,
							"z":      z,
							"type":   name,
							"nodeid": nodeid,
						}
						log.WithFields(fields).Debug("OnEvent()")

						objs := v.onMapObject(pmb.Pos, x, y, z, pmb.Mapblock)

						if len(objs) > 0 {
							for _, obj := range objs {
								err := l.ctx.Objectdb.AddMapData(obj)
								if err != nil {
									fields = logrus.Fields{
										"mbpos": pmb.Pos,
										"x":     x,
										"y":     y,
										"z":     z,
										"type":  name,
										"obj":   obj,
									}
									log.WithFields(fields).Error("AddMapData()")

									//unrecoverable
									panic(err)
								}

								l.ctx.WebEventbus.Emit("mapobject-created", obj)
							}
						}
					}
				})
			} // k==name
		} //for k,v

		for k, v := range l.objectlisteners {
			if k == name {
				//block matches
				coords.IterateMapblock(func(x, y, z int) {
					nodeid := pmb.Mapblock.GetNodeId(x, y, z)
					if nodeid == id {
						fields := logrus.Fields{
							"mbpos":  pmb.Pos,
							"x":      x,
							"y":      y,
							"z":      z,
							"type":   name,
							"nodeid": nodeid,
						}
						log.WithFields(fields).Debug("OnEvent()")

						obj := v.onMapObject(pmb.Pos, x, y, z, pmb.Mapblock)

						if obj != nil {
							err := l.ctx.Objectdb.AddMapData(obj)
							if err != nil {
								fields = logrus.Fields{
									"mbpos": pmb.Pos,
									"x":     x,
									"y":     y,
									"z":     z,
									"type":  name,
									"obj":   obj,
								}
								log.WithFields(fields).Error("AddMapData()")

								//unrecoverable
								panic(err)
							}
							l.ctx.WebEventbus.Emit("mapobject-created", obj)
						}
					}
				})
			} // k==name
		} //for k,v

	} //for id, name
}
