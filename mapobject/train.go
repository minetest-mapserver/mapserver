package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type TrainBlock struct{}

func (this *TrainBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "train")
	o.Attributes["station"] = md["station"]
	o.Attributes["line"] = md["line"]
	o.Attributes["index"] = md["index"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["color"] = md["color"]
	o.Attributes["rail_pos"] = md["rail_pos"]
	o.Attributes["linepath_from_prv"] = md["linepath_from_prv"]

	return o
}
