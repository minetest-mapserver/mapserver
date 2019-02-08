package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type LuaControllerBlock struct{}

func (this *LuaControllerBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "luacontroller")

	//TODO: is this private?
	o.Attributes["code"] = md["code"]

	return o
}
