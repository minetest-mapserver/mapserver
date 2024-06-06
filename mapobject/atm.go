package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

type ATM struct{}

func (this *ATM) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	nodename := block.GetNodeName(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "atm")

	switch nodename {
		case "atm:wtt", "um_wtt:wtt":
			o.Attributes["type"] = "wiretransfer"
		case "atm:atm2", "um_atm:atm_2":
			o.Attributes["type"] = "atm2"
		case "atm:atm3", "um_atm:atm_3":
			o.Attributes["type"] = "atm3"
		default:
			o.Attributes["type"] = "atm"
	}

	return o
}
