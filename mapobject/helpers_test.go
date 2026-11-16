package mapobject

import (
	"fmt"
	"mapserver/types"

	mtypes "github.com/luanti-golang/types"
)

var testMBPos = types.NewMapBlockCoords(1, -2, 3)

// newTestBlock creates a mapblock with a single node of the given name at (x,y,z)
func newTestBlock(nodename string, x, y, z int) *mtypes.MapBlock {
	mb := mtypes.NewMapblock()
	mb.Size = 4096
	mb.ContentId = make([]int, 4096)
	mb.Param1 = make([]int, 4096)
	mb.Param2 = make([]int, 4096)
	mb.BlockMapping[0] = "air"
	mb.BlockMapping[1] = nodename
	mb.ContentId[mtypes.NewPos(x, y, z).Index()] = 1
	return mb
}

func setMeta(mb *mtypes.MapBlock, x, y, z int, pairs map[string]string) {
	i := mtypes.NewPos(x, y, z).Index()
	if mb.Fields[i] == nil {
		mb.Fields[i] = map[string]string{}
	}
	for k, v := range pairs {
		mb.Fields[i][k] = v
	}
}

// sets the inventory list in the serialized form the parser provides
func setInv(mb *mtypes.MapBlock, x, y, z int, name string, items ...*Item) {
	i := mtypes.NewPos(x, y, z).Index()
	if mb.Inventory[i] == nil {
		mb.Inventory[i] = map[string][]string{}
	}
	slots := make([]string, len(items))
	for n, item := range items {
		if !item.IsEmpty() {
			slots[n] = fmt.Sprintf("%s %d", item.Name, item.Count)
		}
	}
	mb.Inventory[i][name] = slots
}
