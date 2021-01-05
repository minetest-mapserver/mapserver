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
	if ctx.Config.MapObjects.MapserverPOI {
		l.AddMapObject("mapserver:poi", &PoiBlock{Color: "blue"})
		l.AddMapObject("mapserver:poi_blue", &PoiBlock{Color: "blue"})
		l.AddMapObject("mapserver:poi_green", &PoiBlock{Color: "green"})
		l.AddMapObject("mapserver:poi_orange", &PoiBlock{Color: "orange"})
		l.AddMapObject("mapserver:poi_red", &PoiBlock{Color: "red"})
		l.AddMapObject("mapserver:poi_purple", &PoiBlock{Color: "purple"})
	}

	if ctx.Config.MapObjects.MapserverTrainline {
		l.AddMapObject("mapserver:train", &TrainBlock{})
	}

	if ctx.Config.MapObjects.MapserverBorder {
		l.AddMapObject("mapserver:border", &BorderBlock{})
	}

	if ctx.Config.MapObjects.MapserverLabel {
		l.AddMapObject("mapserver:label", &LabelBlock{})
	}

	//old tileserver stuff
	if ctx.Config.MapObjects.TileServerLegacy {
		l.AddMapObject("tileserver:poi", &PoiBlock{})
		l.AddMapObject("tileserver:train", &TrainBlock{})
	}

	//travelnet
	if ctx.Config.MapObjects.Travelnet {
		l.AddMapObject("travelnet:travelnet", &TravelnetBlock{})
	}

	//protector
	if ctx.Config.MapObjects.Protector {
		l.AddMapObject("protector:protect", &ProtectorBlock{})
		l.AddMapObject("protector:protect2", &ProtectorBlock{})
	}

	//xp protector
	if ctx.Config.MapObjects.XPProtector {
		l.AddMapObject("xp_redo:protector", &XPProtectorBlock{})
	}

	//priv protector
	if ctx.Config.MapObjects.PrivProtector {
		l.AddMapObject("priv_protector:protector", &PrivProtectorBlock{})
	}

	//builtin
	if ctx.Config.MapObjects.Bones {
		l.AddMapObject("bones:bones", &BonesBlock{})
	}

	//technic
	if ctx.Config.MapObjects.TechnicQuarry {
		l.AddMapObject("technic:quarry", &QuarryBlock{})
	}

	if ctx.Config.MapObjects.TechnicReactor {
		l.AddMapObject("technic:hv_nuclear_reactor_core_active", &NuclearReactorBlock{})
	}

	if ctx.Config.MapObjects.TechnicAnchor {
		l.AddMapObject("technic:admin_anchor", &TechnicAnchorBlock{})
	}

	if ctx.Config.MapObjects.TechnicSwitch {
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
		v := &FancyVend{}
		l.AddMapObject("fancy_vend:admin_vendor", v)
		l.AddMapObject("fancy_vend:admin_depo", v)
		l.AddMapObject("fancy_vend:player_vendor", v)
		l.AddMapObject("fancy_vend:player_depo", v)
	}

	if ctx.Config.MapObjects.ATM {
		atm := &ATM{}
		l.AddMapObject("atm:atm", atm)
		l.AddMapObject("atm:atm2", atm)
		l.AddMapObject("atm:atm3", atm)
		l.AddMapObject("atm:wtt", atm)
	}

	//locator
	if ctx.Config.MapObjects.Locator {
		loc := &Locator{}
		l.AddMapObject("locator:beacon_1", loc)
		l.AddMapObject("locator:beacon_2", loc)
		l.AddMapObject("locator:beacon_3", loc)
	}

	//signs
	if ctx.Config.MapObjects.Signs {
		l.AddMapObject("default:sign_wall_wood", &SignBlock{Material: "wood"})
		l.AddMapObject("default:sign_wall_steel", &SignBlock{Material: "steel"})
	}

	ctx.MapBlockAccessor.Eventbus.AddListener(&l)
}
