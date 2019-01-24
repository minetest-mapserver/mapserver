package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

func onPoiBlock(id int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor) {

	for x:=0; x<16; x++ {
		for y:=0; y<16; y++ {
			for z:=0; z<16; z++ {
				name := block.GetNodeName(x,y,z)
				if name == "mapserver:poi" {
					o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "poi")
					o.Attributes["name"] = "test"

					odb.AddMapData(o)
				}
			}
		}
	}

	panic("OK") //XXX
}
