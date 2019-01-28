package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type ProtectorBlock struct {}

func (this *ProtectorBlock) onMapObject(x,y,z int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor) {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "protector")
	o.Attributes["owner"] = md["owner"]

	odb.AddMapData(o)
}
