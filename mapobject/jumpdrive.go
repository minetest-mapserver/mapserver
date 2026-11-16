package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type JumpdriveBlock struct{}

func (blk *JumpdriveBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "jumpdrive")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["radius"] = md["radius"]

	return o
}
