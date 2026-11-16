package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type XPProtectorBlock struct{}

func (blk *XPProtectorBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "xpprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["xpthreshold"] = md["xpthreshold"]

	return o
}
