package postgres

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var migrations embed.FS

type PostgresAccessor struct {
	db *sql.DB
}

func (a *PostgresAccessor) Migrate() error {
	driver, err := postgres.WithInstance(a.db, &postgres.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
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
