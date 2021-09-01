package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
	"strings"
)

type TravelnetBlock struct{}

func (tn *TravelnetBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	// ignore (P) prefixed stations
	// TODO: configurable prefix
	if strings.HasPrefix(md["station_name"], "(P)") {
		return nil
	}

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "travelnet")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["station_name"] = md["station_name"]
	o.Attributes["station_network"] = md["station_network"]
	o.Attributes["nodename"] = block.GetNodeName(x, y, z)

	return o
}
