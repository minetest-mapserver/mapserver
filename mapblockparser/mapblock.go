package mapblockparser


type MapBlock struct {
	Version byte
	Underground bool
	Mapdata []byte
	Metadata Metadata
}