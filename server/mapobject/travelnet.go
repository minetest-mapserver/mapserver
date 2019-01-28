package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type TravelnetBlock struct {}

func (this *TravelnetBlock) onMapObject(x,y,z int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor) {
	md := block.Metadata.GetMetadata(x, y, z)

	if md["station_name"] == "" || md["owner"] == "" {
		//station not set up
		return
	}

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "travelnet")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["station_name"] = md["station_name"]
	o.Attributes["station_network"] = md["station_network"]

	odb.AddMapData(o)
}
