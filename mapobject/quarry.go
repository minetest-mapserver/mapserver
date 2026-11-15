package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type QuarryBlock struct{}

func (blk *QuarryBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	if md["owner"] == "" {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "technicquarry")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["dug"] = md["dug"]
	o.Attributes["enabled"] = md["enabled"]

	return o
}
