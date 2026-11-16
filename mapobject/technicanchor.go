package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type TechnicAnchorBlock struct{}

func (blk *TechnicAnchorBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicanchor")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]
	o.Attributes["locked"] = md["locked"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
