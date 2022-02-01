package sqlite

import (
	"database/sql"
	"mapserver/public"
	"time"

	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {
	log.WithFields(logrus.Fields{"filename": db.filename}).Info("Migrating database")
	start := time.Now()

	sql, err := public.Files.ReadFile("sql/sqlite_mapobjectdb_migrate.sql")
	if err != nil {
		return err
	}

	_, err = db.db.Exec(string(sql))
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
