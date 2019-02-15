package mapobject

import (
	"mapserver/app"
)

func Setup(ctx *app.App) {
	l := Listener{
		ctx:                  ctx,
		objectlisteners:      make(map[string]MapObjectListener),
		multiobjectlisteners: make(map[string]MapMultiObjectListener),
	}

	//mapserver stuff
	if ctx.Config.MapObjects.Mapserver {
		l.AddMapObject("mapserver:poi", &PoiBlock{})
		l.AddMapObject("mapserver:train", &TrainBlock{})
		l.AddMapObject("mapserver:border", &BorderBlock{})
		l.AddMapObject("mapserver:label", &LabelBlock{})
	}

	//travelnet
	if ctx.Config.MapObjects.Travelnet {
		l.AddMapObject("travelnet:travelnet", &TravelnetBlock{})
	}

	//protections
	if ctx.Config.MapObjects.Protector {
		l.AddMapObject("protector:protect", &ProtectorBlock{})
		l.AddMapObject("protector:protect2", &ProtectorBlock{})
		l.AddMapObject("xp_redo:protector", &XPProtectorBlock{})
	}

	//builtin
	if ctx.Config.MapObjects.Bones {
		l.AddMapObject("bones:bones", &BonesBlock{})
	}

	//technic
	if ctx.Config.MapObjects.Technic {
		l.AddMapObject("technic:quarry", &QuarryBlock{})
		l.AddMapObject("technic:hv_nuclear_reactor_core_active", &NuclearReactorBlock{})
		l.AddMapObject("technic:admin_anchor", &TechnicAnchorBlock{})
		l.AddMapObject("technic:switching_station", &TechnicSwitchBlock{})
	}

	//digilines
	if ctx.Config.MapObjects.Digilines {
		l.AddMapObject("digilines:lcd", &DigilineLcdBlock{})
	}

	//mesecons
	if ctx.Config.MapObjects.LuaController {
		luac := &LuaControllerBlock{}
		// mesecons_luacontroller:luacontroller0000 2^4=16
		l.AddMapObject("mesecons_luacontroller:luacontroller1111", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1110", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1100", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1010", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1000", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1101", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1001", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller1011", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0111", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0110", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0100", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0010", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0000", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0101", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0001", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller0011", luac)
		l.AddMapObject("mesecons_luacontroller:luacontroller_burnt", luac)
	}

	//digiterms
	if ctx.Config.MapObjects.Digiterms {
		digiterms := &DigitermsBlock{}
		l.AddMapObject("digiterms:lcd_monitor", digiterms)
		l.AddMapObject("digiterms:cathodic_beige_monitor", digiterms)
		l.AddMapObject("digiterms:cathodic_white_monitor", digiterms)
		l.AddMapObject("digiterms:cathodic_black_monitor", digiterms)
		l.AddMapObject("digiterms:scifi_glassscreen", digiterms)
		l.AddMapObject("digiterms:scifi_widescreen", digiterms)
		l.AddMapObject("digiterms:scifi_tallscreen", digiterms)
		l.AddMapObject("digiterms:scifi_keysmonitor", digiterms)
	}

	//missions
	if ctx.Config.MapObjects.Mission {
		l.AddMapObject("missions:mission", &MissionBlock{})
	}

	//jumpdrive, TODO: fleet controller
	if ctx.Config.MapObjects.Jumpdrive {
		l.AddMapObject("jumpdrive:engine", &JumpdriveBlock{})
	}

	//smartshop
	if ctx.Config.MapObjects.Smartshop {
		l.AddMapMultiObject("smartshop:shop", &SmartShopBlock{})
	}

	if ctx.Config.MapObjects.Fancyvend {
		//TODO
	}

	if ctx.Config.MapObjects.ATM {
		//TODO
	}

	ctx.BlockAccessor.Eventbus.AddListener(&l)
}
