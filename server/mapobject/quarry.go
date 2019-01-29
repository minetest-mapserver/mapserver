package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type QuarryBlock struct{}

func (this *QuarryBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["owner"] == "" {
		return nil
	}

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "technicquarry")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["dug"] = md["dug"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
