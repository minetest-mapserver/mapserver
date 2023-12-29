package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type LabelBlock struct{}

func (this *LabelBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "label")
	o.Attributes["text"] = md["text"]
	o.Attributes["size"] = md["size"]
	o.Attributes["direction"] = md["direction"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["color"] = md["color"]

	return o
}
