package sqlite

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrations embed.FS

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (a *Sqlite3Accessor) Migrate() error {
	driver, err := sqlite3.WithInstance(a.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "sqlite", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	v, _, _ := m.Version()
	logrus.WithFields(logrus.Fields{"version": v}).Info("DB Migrated")

	return nil
}

func (a *Sqlite3Accessor) EnableSpeedSafetyTradeoff(enableSpeed bool) error {
	if enableSpeed {
		_, err := a.db.Exec("PRAGMA journal_mode = MEMORY; PRAGMA synchronous = OFF;")
		return err

	} else {
		_, err := a.db.Exec("PRAGMA journal_mode = TRUNCATE; PRAGMA synchronous = ON;")
		return err

	}
}
