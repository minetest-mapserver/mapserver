package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	mtypes "github.com/minetest-go/types"
)

type UnifiedMoneyAreaForSale struct{}

func (um *UnifiedMoneyAreaForSale) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	if _, ok := md["id"]; !ok {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "um_area_forsale")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["id"] = md["id"] // ", " seperated
	o.Attributes["price"] = md["price"]
	o.Attributes["description"] = md["description"]

	return o
}
