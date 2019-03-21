package postgres

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func New(connStr string) (*PostgresAccessor, error) {
	db, err := sql.Open("postgres", connStr+" sslmode=disable")

	if err != nil {
		return nil, err
	}

	sq := &PostgresAccessor{db: db}
	return sq, nil
}
