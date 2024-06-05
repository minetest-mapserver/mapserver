package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type UnifiefMoneyAreaForSale struct{}

func (this *UnifiefMoneyAreaForSale) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "um_area_forsale")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["id"] = md["id"] // ", " seperated
	o.Attributes["price"] = md["price"]
	o.Attributes["description"] = md["description"]

	return o
}
