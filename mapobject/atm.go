package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type ATM struct{}

func (this *ATM) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	nodename := block.GetNodeName(x, y, z)

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "atm")

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
