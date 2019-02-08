package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type DigitermsBlock struct{}

func (this *DigitermsBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "digiterm")
	o.Attributes["display_text"] = md["display_text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
