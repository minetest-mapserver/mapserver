package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type PoiBlock struct {}

func (this *PoiBlock) onMapObject(x,y,z int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor) {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "poi")
	o.Attributes["name"] = md["name"]
	o.Attributes["category"] = md["category"]
	o.Attributes["url"] = md["url"]
	o.Attributes["active"] = md["active"]
	o.Attributes["owner"] = md["owner"]

	odb.AddMapData(o)
}
