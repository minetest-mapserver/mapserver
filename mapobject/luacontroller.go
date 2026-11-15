package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type LuaControllerBlock struct{}

func (l *LuaControllerBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	//md := getMetadata(block, x, y, z)
	nodename := block.GetNodeName(mtypes.NewPos(x, y, z))

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
