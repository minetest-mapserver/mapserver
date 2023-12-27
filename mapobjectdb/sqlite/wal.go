package sqlite

import (
	"database/sql"
	"errors"
)

func EnableWAL(db *sql.DB) error {
	result := db.QueryRow("pragma journal_mode;")
	var mode string
	err := result.Scan(&mode)
	if err != nil {
		return err
	}

	if mode != "wal" {
		_, err = db.Exec("pragma journal_mode = wal;")
		if err != nil {
			return errors.New("couldn't switch the db-journal to wal-mode, please stop the minetest-engine to allow doing this or do it manually: " + err.Error())
		}
	}

	return nil
}
