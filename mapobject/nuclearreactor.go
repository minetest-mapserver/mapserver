package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

type NuclearReactorBlock struct{}

func (blk *NuclearReactorBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "nuclearreactor")
	o.Attributes["burn_time"] = md["burn_time"]
	o.Attributes["structure_accumulated_badness"] = md["structure_accumulated_badness"]

	return o
}
