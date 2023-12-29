package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type DigitermsBlock struct{}

func (this *DigitermsBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "digiterm")
	o.Attributes["display_text"] = md["display_text"]
	o.Attributes["channel"] = md["channel"]

	return o
}
