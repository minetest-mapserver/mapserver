package mapobjectdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	//TODO: flag/config for unsafe db access
	db, err := sql.Open("sqlite3", filename+"?_timeout=500&_journal_mode=MEMORY&_synchronous=OFF")
	db.SetMaxOpenConns(1)

	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
