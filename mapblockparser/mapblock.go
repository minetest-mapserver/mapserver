package mapblockparser


type MapBlock struct {
	Version byte
	Underground bool
	Mapdata []byte
	Metadata Metadata
}


type Metadata struct {
	Inventories map[int]map[string]Inventory
	Pairs map[int]map[string]string
}

type Item struct {
	Name string
	Count int
	Wear int
}

type Inventory struct {
	Size int
	Items []Item
}