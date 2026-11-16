package mapobject

import (
	"mapserver/mapobjectdb"
	"mapserver/types"
	"strings"

	mtypes "github.com/luanti-golang/types"
)

type TravelnetBlock struct{}

func (tn *TravelnetBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mtypes.MapBlock) *mapobjectdb.MapObject {
	md := getMetadata(block, x, y, z)

	// ignore (P) prefixed stations
	// TODO: configurable prefix
	if strings.HasPrefix(md["station_name"], "(P)") {
		return nil
	}

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "travelnet")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["station_name"] = md["station_name"]
	o.Attributes["station_network"] = md["station_network"]
	o.Attributes["nodename"] = block.GetNodeName(mtypes.NewPos(x, y, z))

	return o
}
