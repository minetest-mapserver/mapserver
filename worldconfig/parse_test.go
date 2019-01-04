package worldconfig_test

import (
	"testing"
	worldconfig "mapserver/worldconfig"
)

func TestParseSqlite(t *testing.T) {
	cfg := worldconfig.Parse("./testdata/world.mt.sqlite")
	if cfg.Backend != worldconfig.BACKEND_SQLITE3 {
		t.Fatal("not sqlite3")
	}
	if cfg.PlayerBackend != worldconfig.BACKEND_FILES {
		t.Fatal("not files")
	}
}

func TestParsePostgres(t *testing.T) {
	cfg := worldconfig.Parse("./testdata/world.mt.postgres")
	if cfg.Backend != worldconfig.BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg.PlayerBackend != worldconfig.BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg.PsqlConnection == nil {
		t.Fatal("no connection")
	}

	if cfg.PsqlPlayerConnection == nil {
		t.Fatal("no connection")
	}

	if cfg.PsqlConnection.Host != "postgres" {
		t.Fatal("param err")
	}

	if cfg.PsqlConnection.Port != 5432 {
		t.Fatal("param err: ", cfg.PsqlConnection.Port)
	}
}
