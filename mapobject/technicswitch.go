package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type TechnicSwitchBlock struct{}

func (blk *TechnicSwitchBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicswitch")
	o.Attributes["active"] = md["active"]
	o.Attributes["channel"] = md["channel"]
	o.Attributes["supply"] = md["supply"]
	o.Attributes["demand"] = md["demand"]

	return o
}
