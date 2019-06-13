package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type BorderBlock struct{}

func (this *BorderBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "border")
	o.Attributes["name"] = md["name"]
	o.Attributes["index"] = md["index"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["color"] = md["color"]

	return o
}
