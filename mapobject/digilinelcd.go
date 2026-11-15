package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type DigilineLcdBlock struct{}

func (blk *DigilineLcdBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "digilinelcd")
	o.Attributes["text"] = md["text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
