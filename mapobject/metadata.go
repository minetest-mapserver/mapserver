package mapobject

import (
	"strconv"
	"strings"

	mtypes "github.com/luanti-golang/types"
)

// Item is a parsed itemstring of an inventory slot
type Item struct {
	Name  string
	Count int
	Wear  int
}

func (i *Item) IsEmpty() bool {
	return i.Name == "" && i.Count == 0
}

// Inventory holds the non-empty slots of an inventory list
type Inventory struct {
	Items []*Item
}

func (inv *Inventory) IsEmpty() bool {
	for _, item := range inv.Items {
		if item.Name != "" && item.Count > 0 {
			return false
		}
	}

	return true
}

// parses an itemstring in the form "name [count [wear [metadata]]]"
func parseItem(str string) *Item {
	parts := strings.Split(str, " ")
	item := &Item{Name: parts[0], Count: 1}

	if len(parts) >= 2 {
		if val, err := strconv.Atoi(parts[1]); err == nil {
			item.Count = val
		}
	}

	if len(parts) >= 3 {
		if val, err := strconv.Atoi(parts[2]); err == nil {
			item.Wear = val
		}
	}

	return item
}

// returns the metadata fields of the node, the result is nil if there are none
func getMetadata(block *mtypes.MapBlock, x, y, z int) map[string]string {
	return block.Fields[mtypes.NewPos(x, y, z).Index()]
}

// returns the parsed inventories of the node by list name, empty slots are skipped
func getInventoryMap(block *mtypes.MapBlock, x, y, z int) map[string]*Inventory {
	invMap := map[string]*Inventory{}

	for name, slots := range block.Inventory[mtypes.NewPos(x, y, z).Index()] {
		inv := &Inventory{}
		for _, slot := range slots {
			if slot != "" {
				inv.Items = append(inv.Items, parseItem(slot))
			}
		}
		invMap[name] = inv
	}

	return invMap
}
