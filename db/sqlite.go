package db

type Sqlite3Accessor struct {
}

func (db *Sqlite3Accessor) IsMigrated() (bool, error) {
	return false, nil
}

func (db *Sqlite3Accessor) Migrate() error {
	return nil
}

func (db *Sqlite3Accessor) FindLatestBlocks(mintime int64, limit int) ([]Block, error) {
	return make([]Block, 0)
}

func (db *Sqlite3Accessor) FindBlocks(posx int, posz int, posystart int, posyend int) ([]Block, error) {
	return make([]Block, 0)
}

func (db *Sqlite3Accessor) CountBlocks(x1, x2, y1, y2, z1, z2 int) (int, error) {
	return 0
}

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	return nil, nil
}
