package worldconfig

import (
	"fmt"
	"testing"
)

func TestParseSqlite(t *testing.T) {
	cfg := Parse("./testdata/world.mt.sqlite")
	if cfg.Backend != BACKEND_SQLITE3 {
		t.Fatal("not sqlite3")
	}
}

func TestParsePostgres(t *testing.T) {
	cfg := Parse("./testdata/world.mt.postgres")
	fmt.Println(cfg)
	if cfg.Backend != BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg.PsqlConnection != "host=postgres port=5432 user=postgres password=enter dbname=postgres" {
		t.Fatal("param err")
	}
}
