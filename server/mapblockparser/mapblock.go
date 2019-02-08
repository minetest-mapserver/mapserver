package mapblockparser

import (
	"mapserver/coords"
)

type MapBlock struct {
	Pos          *coords.MapBlockCoords `json:"pos"`
	Size         int                    `json:"size"`
	Version      byte                   `json:"version"`
	Underground  bool                   `json:"underground"`
	Mapdata      *MapData               `json:"mapdata"`
	Metadata     *Metadata              `json:"metadata"`
	BlockMapping map[int]string         `json:"blockmapping"`
	Mtime        int64                  `json:"mtime"`
}

type MapData struct {
	ContentId []int `json:"contentid"`
	Param1    []int `json:"param1"`
	Param2    []int `json:"param2"`
}

type Metadata struct {
	Inventories map[int]map[string]*Inventory `json:"inventories"`
	Pairs       map[int]map[string]string     `json:"pairs"`
}

type Item struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Wear  int    `json:"wear"`
}

type Inventory struct {
	Size  int    `json:"size"`
	Items []Item `json:"items"`
}

func getNodePos(x, y, z int) int {
	return x + (y * 16) + (z * 256)
}

func (inv *Inventory) IsEmpty() bool {
	if inv.Size == 0 || len(inv.Items) == 0 {
		return true
	}

	for _, item := range inv.Items {
		if item.Name != "" && item.Count > 0 {
			return false
		}
	}

	return true
}

func (mb *MapBlock) GetNodeId(x, y, z int) int {
	pos := getNodePos(x, y, z)
	return mb.Mapdata.ContentId[pos]
}

func (mb *MapBlock) GetNodeName(x, y, z int) string {
	id := mb.GetNodeId(x, y, z)
	return mb.BlockMapping[id]
}

func NewMapblock() *MapBlock {
	mb := MapBlock{}
	mb.Metadata = NewMetadata()
	mb.BlockMapping = make(map[int]string)
	return &mb
}

func NewMetadata() *Metadata {
	md := Metadata{}
	md.Inventories = make(map[int]map[string]*Inventory)
	md.Pairs = make(map[int]map[string]string)
	return &md
}

func (md *Metadata) GetMetadata(x, y, z int) map[string]string {
	return md.GetPairsMap(getNodePos(x, y, z))
}

func (md *Metadata) GetPairsMap(pos int) map[string]string {
	pairsMap := md.Pairs[pos]
	if pairsMap == nil {
		pairsMap = make(map[string]string)
		md.Pairs[pos] = pairsMap
	}

	return pairsMap
}

func (md *Metadata) GetInventoryMap(pos int) map[string]*Inventory {
	invMap := md.Inventories[pos]
	if invMap == nil {
		invMap = make(map[string]*Inventory)
		md.Inventories[pos] = invMap
	}

	return invMap
}

func (md *Metadata) GetInventoryMapAtPos(x, y, z int) map[string]*Inventory {
	return md.GetInventoryMap(getNodePos(x, y, z))
}

func (md *Metadata) GetInventory(pos int, name string) *Inventory {
	m := md.GetInventoryMap(pos)
	inv := m[name]
	if inv == nil {
		inv = &Inventory{}
		m[name] = inv
	}

	return inv
}
