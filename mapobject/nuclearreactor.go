package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type NuclearReactorBlock struct{}

func (this *NuclearReactorBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "nuclearreactor")
	o.Attributes["burn_time"] = md["burn_time"]
	o.Attributes["structure_accumulated_badness"] = md["structure_accumulated_badness"]

	return o
}
