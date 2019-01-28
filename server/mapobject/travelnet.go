package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type TravelnetBlock struct{}

func (this *TravelnetBlock) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["station_name"] == "" || md["owner"] == "" {
		//station not set up
		return nil
	}

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "travelnet")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["station_name"] = md["station_name"]
	o.Attributes["station_network"] = md["station_network"]

	return o
}
