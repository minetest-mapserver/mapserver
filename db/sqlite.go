package db

type Sqlite3Accessor struct {
}

func (db *Sqlite3Accessor) IsMigrated() (bool, error) {
	return false, nil
}

func (db *Sqlite3Accessor) Migrate() error {
	return nil
}

func (db *Sqlite3Accessor) FindLatestBlocks(time int64, limit int) []Block {
	return make([]Block, 0)
}

func (db *Sqlite3Accessor) FindBlocks(posx int, posz int, posystart int, posyend int) []Block {
	return make([]Block, 0)
}

func (db *Sqlite3Accessor) GetXRange(posystart int, posyend int) (int, int) {
	return 0, 0
}

func (db *Sqlite3Accessor) GetZRange(posystart int, posyend int) (int, int) {
	return 0, 0
}

func NewSqliteAccessor(filename string) {

}
