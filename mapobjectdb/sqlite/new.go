package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", filename+"?_timeout=500")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
