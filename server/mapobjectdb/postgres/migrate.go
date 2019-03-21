package postgres

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"mapserver/vfs"
	"time"
)

type PostgresAccessor struct {
	db *sql.DB
}

func (db *PostgresAccessor) Migrate() error {
	log.Info("Migrating database")
	start := time.Now()
	_, err := db.db.Exec(vfs.FSMustString(false, "/sql/postgres_mapobjectdb_migrate.sql"))
	if err != nil {
		return err
	}
	t := time.Now()
	elapsed := t.Sub(start)
	log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")

	return nil
}
