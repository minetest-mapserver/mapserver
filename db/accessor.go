package db

type Block struct {
	posx, posy, posz int
	data             []byte
	mtime            int64
}

type DBAccessor interface {
	IsMigrated() (bool, error)
	Migrate() error
	FindLatestBlocks(time int64, limit int) []Block

	//Block range lookup
	FindBlocks(posx int, posz int, posystart int, posyend int) []Block
	GetXRange(posystart int, posyend int) (int, int)
	GetZRange(posystart int, posyend int) (int, int)
}
