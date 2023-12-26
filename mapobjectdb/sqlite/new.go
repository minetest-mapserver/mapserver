package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", filename+"?busy_timeout=5000")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)

	err = EnableWAL(db)
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db}
	return sq, nil
}
