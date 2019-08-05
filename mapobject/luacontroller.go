package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type LuaControllerBlock struct{}

func (this *LuaControllerBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	//md := block.Metadata.GetMetadata(x, y, z)
	nodename := block.GetNodeName(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "luacontroller")
	//o.Attributes["code"] = md["code"]
	//o.Attributes["lc_memory"] = md["lc_memory"]

	if nodename == "mesecons_luacontroller:luacontroller_burnt" {
		o.Attributes["burnt"] = "1"
	} else {
		o.Attributes["burnt"] = "0"
	}

	return o
}
