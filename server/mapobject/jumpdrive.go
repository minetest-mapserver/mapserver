package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type JumpdriveBlock struct{}

func (this *JumpdriveBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "jumpdrive")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]

	return o
}
