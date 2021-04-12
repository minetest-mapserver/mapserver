package postgres

import (
	"database/sql"
	"mapserver/public"

	"github.com/sirupsen/logrus"

	"time"
)

type PostgresAccessor struct {
	db *sql.DB
}

func (db *PostgresAccessor) Migrate() error {
	log.Info("Migrating database")
	start := time.Now()
	sql, err := public.Files.ReadFile("sql/postgres_mapobjectdb_migrate.sql")
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
