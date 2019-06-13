package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type PrivProtectorBlock struct{}

func (this *PrivProtectorBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "privprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["priv"] = md["priv"]

	return o
}
