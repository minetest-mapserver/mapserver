package mapobject

import (
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
)

type TrainBlock struct {}

func (this *TrainBlock) onMapObject(x,y,z int, block *mapblockparser.MapBlock, odb mapobjectdb.DBAccessor) {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(&block.Pos, x, y, z, "train")
	o.Attributes["station"] = md["station"]
	o.Attributes["line"] = md["line"]
	o.Attributes["index"] = md["index"]
	o.Attributes["owner"] = md["owner"]

	odb.AddMapData(o)
}
