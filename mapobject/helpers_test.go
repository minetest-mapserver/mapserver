package mapobject

import (
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

var testMBPos = types.NewMapBlockCoords(1, -2, 3)

// newTestBlock creates a mapblock with a single node of the given name at (x,y,z)
func newTestBlock(nodename string, x, y, z int) *mapparser.MapBlock {
	mb := mapparser.NewMapblock()
	mb.Size = 4096
	mb.Mapdata = &mapparser.MapData{
		ContentId: make([]int, 4096),
		Param1:    make([]int, 4096),
		Param2:    make([]int, 4096),
	}
	mb.BlockMapping[0] = "air"
	mb.BlockMapping[1] = nodename
	mb.Mapdata.ContentId[mapparser.GetNodePos(x, y, z)] = 1
	return mb
}

func setMeta(mb *mapparser.MapBlock, x, y, z int, pairs map[string]string) {
	md := mb.Metadata.GetMetadata(x, y, z)
	for k, v := range pairs {
		md[k] = v
	}
}

func setInv(mb *mapparser.MapBlock, x, y, z int, name string, items ...*mapparser.Item) {
	mb.Metadata.GetInventoryMapAtPos(x, y, z)[name] = &mapparser.Inventory{Size: len(items), Items: items}
}
