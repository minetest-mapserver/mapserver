package worldconfig

import (
	"testing"
	"fmt"
)

func TestParseSqlite(t *testing.T) {
	cfg := Parse("./testdata/world.mt.sqlite")
	if cfg.Backend != BACKEND_SQLITE3 {
		t.Fatal("not sqlite3")
	}
	if cfg.PlayerBackend != BACKEND_FILES {
		t.Fatal("not files")
	}
}

func TestParsePostgres(t *testing.T) {
	cfg := Parse("./testdata/world.mt.postgres")
	fmt.Println(cfg)
	if cfg.Backend != BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg.PlayerBackend != BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg.PsqlConnection.Host != "postgres" {
		t.Fatal("param err")
	}

	if cfg.PsqlConnection.Port != 5432 {
		t.Fatal("param err: ", cfg.PsqlConnection.Port)
	}
}
