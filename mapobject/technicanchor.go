package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type TechnicAnchorBlock struct{}

func (this *TechnicAnchorBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicanchor")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]
	o.Attributes["locked"] = md["locked"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
