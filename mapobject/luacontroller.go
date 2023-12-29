package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type LuaControllerBlock struct{}

func (this *LuaControllerBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	//md := block.Metadata.GetMetadata(x, y, z)
	nodename := block.GetNodeName(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "luacontroller")
	//o.Attributes["code"] = md["code"]
	//o.Attributes["lc_memory"] = md["lc_memory"]

	if nodename == "mesecons_luacontroller:luacontroller_burnt" {
		o.Attributes["burnt"] = "1"
	} else {
		o.Attributes["burnt"] = "0"
	}

	return o
}
