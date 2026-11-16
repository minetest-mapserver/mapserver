package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type PrivProtectorBlock struct{}

func (blk *PrivProtectorBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "privprotector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["priv"] = md["priv"]

	return o
}
