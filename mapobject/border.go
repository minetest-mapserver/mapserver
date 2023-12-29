package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type BorderBlock struct{}

func (this *BorderBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "border")
	o.Attributes["name"] = md["name"]
	o.Attributes["index"] = md["index"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["color"] = md["color"]

	return o
}
