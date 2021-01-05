package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type SignBlock struct {
	Material string
}

func (this *SignBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "sign")
	o.Attributes["display_text"] = md["text"]
	o.Attributes["material"] = this.Material

	return o
}
