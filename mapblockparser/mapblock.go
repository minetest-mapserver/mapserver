package mapblockparser

type MapBlock struct {
	Version      byte
	Underground  bool
	Mapdata      []byte
	Metadata     Metadata
	BlockMapping map[int]string
}

func getNodePos(x, y, z int) int {
	return x + (y * 16) + (z * 256)
}

func (mb *MapBlock) GetNodeName(x, y, z int) string {
	pos := getNodePos(x, y, z)
	id := readU16(mb.Mapdata, pos*2)
	return mb.BlockMapping[id]
}

func NewMapblock() MapBlock {
	mb := MapBlock{}
	mb.Metadata = NewMetadata()
	mb.BlockMapping = make(map[int]string)
	return mb
}

type Metadata struct {
	Inventories map[int]map[string]*Inventory
	Pairs       map[int]map[string]string
}

func NewMetadata() Metadata {
	md := Metadata{}
	md.Inventories = make(map[int]map[string]*Inventory)
	md.Pairs = make(map[int]map[string]string)
	return md
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

func (md *Metadata) GetInventory(pos int, name string) *Inventory {
	m := md.GetInventoryMap(pos)
	inv := m[name]
	if inv == nil {
		inv = &Inventory{}
		m[name] = inv
	}

	return inv
}

type Item struct {
	Name  string
	Count int
	Wear  int
}

type Inventory struct {
	Size  int
	Items []Item
}
