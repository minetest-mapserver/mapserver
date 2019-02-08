package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type LuaControllerBlock struct{}

func (this *LuaControllerBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "luacontroller")
	o.Attributes["code"] = md["code"]
	o.Attributes["lc_memory"] = md["lc_memory"]

	return o
}
