package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type NuclearReactorBlock struct{}

func (this *NuclearReactorBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "nuclearreactor")
	o.Attributes["burn_time"] = md["burn_time"]
	o.Attributes["structure_accumulated_badness"] = md["structure_accumulated_badness"]

	return o
}
