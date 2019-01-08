package db

type Block struct {
	posx, posy, posz int
	data             []byte
	mtime            int64
}

type DBAccessor interface {
	IsMigrated() (bool, error)
	Migrate() error
	FindLatestBlocks(mintime int64, limit int) []Block, error
	FindBlocks(posx int, posz int, posystart int, posyend int) []Block, error
	CountBlocks(x1, x2, y1, y2, z1, z2 int) int, error
}
