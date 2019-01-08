package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const migrateScript = `
`

type Sqlite3Accessor struct {
	db *sql.DB
}

func (db *Sqlite3Accessor) Migrate() error {
	return nil
}

func (db *Sqlite3Accessor) FindLatestBlocks(mintime int64, limit int) ([]Block, error) {
	return make([]Block, 0), nil
}

func (db *Sqlite3Accessor) FindBlocks(posx int, posz int, posystart int, posyend int) ([]Block, error) {
	return make([]Block, 0), nil
}

func (db *Sqlite3Accessor) CountBlocks(x1, x2, y1, y2, z1, z2 int) (int, error) {
	return 0, nil
}

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename + "?mode=ro")
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db}
	return sq, nil
}
