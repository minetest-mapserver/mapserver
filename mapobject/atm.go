package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type ATM struct{}

func (this *ATM) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	nodename := block.GetNodeName(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "atm")

	if nodename == "atm:wtt" {
		o.Attributes["type"] = "wiretransfer"

	} else if nodename == "atm:atm2" {
		o.Attributes["type"] = "atm2"

	} else if nodename == "atm:atm3" {
		o.Attributes["type"] = "atm3"

	} else {
		o.Attributes["type"] = "atm"

	}

	return o
}
