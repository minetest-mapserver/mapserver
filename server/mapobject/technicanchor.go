package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type TechnicAnchorBlock struct{}

func (this *TechnicAnchorBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "technicanchor")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]
	o.Attributes["locked"] = md["locked"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
