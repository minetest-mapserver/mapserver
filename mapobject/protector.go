package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type ProtectorBlock struct{}

func (blk *ProtectorBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "protector")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["members"] = md["members"]

	return o
}
