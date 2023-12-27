package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", filename+"?_pragma=busy_timeout(30000)")
	if err != nil {
		return nil, err
	}

	err = EnableWAL(db)
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db}
	return sq, nil
}
