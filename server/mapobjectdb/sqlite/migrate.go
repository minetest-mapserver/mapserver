package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"time"
)

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {
	log.WithFields(logrus.Fields{"filename": db.filename}).Info("Migrating database")
	start := time.Now()
	_, err := db.db.Exec(migrateScript)
	if err != nil {
		return err
	}
	t := time.Now()
	elapsed := t.Sub(start)
	log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")

	return nil
}

func (db *Sqlite3Accessor) EnableSpeedSafetyTradeoff(enableSpeed bool) error {
	if enableSpeed {
		_, err := db.db.Exec("PRAGMA journal_mode = MEMORY; PRAGMA synchronous = OFF;")
		return err

	} else {
		_, err := db.db.Exec("PRAGMA journal_mode = TRUNCATE; PRAGMA synchronous = ON;")
		return err

	}
}
