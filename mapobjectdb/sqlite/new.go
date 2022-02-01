package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", filename+"?_timeout=500")
	db.SetMaxOpenConns(1)

	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
