package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type DigilineLcdBlock struct{}

func (this *DigilineLcdBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "digilinelcd")
	o.Attributes["text"] = md["text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
