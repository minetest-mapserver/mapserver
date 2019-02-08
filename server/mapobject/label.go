package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type LabelBlock struct{}

func (this *LabelBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "label")
	o.Attributes["text"] = md["text"]
	o.Attributes["size"] = md["size"]
	o.Attributes["direction"] = md["direction"]

	return o
}
