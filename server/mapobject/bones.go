package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type BonesBlock struct{}

func (this *BonesBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["owner"] == "" {
		return nil
	}

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "bones")
	o.Attributes["time"] = md["time"]
	o.Attributes["owner"] = md["owner"]

	return o
}
