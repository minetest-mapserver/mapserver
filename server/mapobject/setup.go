package mapobject

import (
	"mapserver/app"
	"mapserver/eventbus"
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type MapObjectListener interface {
	onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject
}

type Listener struct {
	ctx             *app.App
	objectlisteners map[string]MapObjectListener
}

func (this *Listener) AddMapObject(blockname string, ol MapObjectListener) {
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

	this.ctx.WebEventbus.Emit("mapobjects-cleared", block.Pos)

	for id, name := range block.BlockMapping {
		for k, v := range this.objectlisteners {
			if k == name {
				//block matches

				for x := 0; x < 16; x++ {
					for y := 0; y < 16; y++ {
						for z := 0; z < 16; z++ {
							nodeid := block.GetNodeId(x, y, z)
							if nodeid == id {
								obj := v.onMapObject(x, y, z, block)

								if obj != nil {
									this.ctx.Objectdb.AddMapData(obj)
									this.ctx.WebEventbus.Emit("mapobject-created", obj)
								}
							}
						} //z
					} //y
				} //x

			}
		} //for k,v
	} //for id, name
}

func Setup(ctx *app.App) {
	l := Listener{
		ctx:             ctx,
		objectlisteners: make(map[string]MapObjectListener),
	}

	//mapserver stuff
	l.AddMapObject("mapserver:poi", &PoiBlock{})
	l.AddMapObject("mapserver:train", &TrainBlock{})

	//travelnet
	l.AddMapObject("travelnet:travelnet", &TravelnetBlock{})

	//protections
	l.AddMapObject("protector:protect", &ProtectorBlock{})
	l.AddMapObject("protector:protect2", &ProtectorBlock{})
	l.AddMapObject("xp_redo:protector", &XPProtectorBlock{})

	//builtin
	l.AddMapObject("bones:bones", &BonesBlock{})

	//technic
	l.AddMapObject("technic:quarry", &QuarryBlock{})
	l.AddMapObject("technic:hv_nuclear_reactor_core_active", &NuclearReactorBlock{})
	l.AddMapObject("technic:admin_anchor", &TechnicAnchorBlock{})

	//digilines
	l.AddMapObject("digilines:lcd", &DigilineLcdBlock{})

	//missions
	l.AddMapObject("missions:mission", &MissionBlock{})

	//jumpdrive, TODO: fleet controller
	l.AddMapObject("jumpdrive:engine", &JumpdriveBlock{})

	//TODO: atm, digiterms, signs/banners, spacecannons, shops (smart, fancy)

	ctx.BlockAccessor.Eventbus.AddListener(&l)
}
