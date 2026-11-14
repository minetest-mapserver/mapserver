package sqlite

import (
	"database/sql"
	"mapserver/testutils"
	"mapserver/types"
	"mapserver/util"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func TestMigrateEmpty(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateEmpty.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}
}

func TestMigrateEmptyLegacy(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateEmpty.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyLegacyDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}
}

func TestMigrate(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateEmptyDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	err = a.Migrate()
	if err != nil {
		panic(err)
	}
}

func TestMigrateAndQuery(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQuery.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	block, err := a.GetBlock(types.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if block == nil {
		t.Fatal("no data")
	}

}

func TestMigrateAndQueryCount(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	count, err := a.CountBlocks()
	if err != nil {
		panic(err)
	}

	if count <= 0 {
		t.Fatal("zero count")
	}
}

func TestMigrateAndQueryCountLegacy(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabaseLegacy(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	count, err := a.CountBlocks()
	if err != nil {
		panic(err)
	}

	if count <= 0 {
		t.Fatal("zero count")
	}
}

func TestMigrateAndQueryTimestamp(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "TestMigrateAndQueryStride.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	testutils.CreateTestDatabase(tmpfile.Name())
	a, err := New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	count, err := a.GetTimestamp()
	if err != nil {
		panic(err)
	}

	if count <= 0 {
		t.Fatal("zero count")
	}
}

// pages through all blocks with the given limit and returns the visited positions in order
func findAllBlocksByMtime(t *testing.T, a *Sqlite3Accessor, limit int) []types.MapBlockCoords {
	list := []types.MapBlockCoords{}
	seen := map[types.MapBlockCoords]bool{}

	var lastmtime int64 = -1
	var lastpos *types.MapBlockCoords

	for i := 0; i < 100000; i++ {
		blocks, err := a.FindBlocksByMtime(lastmtime, lastpos, limit)
		if err != nil {
			t.Fatal(err)
		}

		if len(blocks) > limit {
			t.Fatalf("got %d blocks with limit %d", len(blocks), limit)
		}

		for _, block := range blocks {
			if block.Mtime < lastmtime {
				t.Fatalf("mtime not ascending: %d after %d", block.Mtime, lastmtime)
			}
			if seen[*block.Pos] {
				t.Fatalf("block returned twice: %v", block.Pos)
			}
			seen[*block.Pos] = true
			list = append(list, *block.Pos)

			lastmtime = block.Mtime
			lastpos = block.Pos
		}

		if len(blocks) < limit {
			return list
		}
	}

	t.Fatal("pagination does not terminate")
	return nil
}

// more blocks with the same mtime than the fetch limit: none of them must be skipped
func testFindBlocksByMtimePagination(t *testing.T, createdb func(string) error) {
	tmpfile, err := os.CreateTemp("", "TestFindBlocksByMtime.*.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if err = createdb(tmpfile.Name()); err != nil {
		t.Fatal(err)
	}

	a, err := New(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// adds the mtime column, all existing blocks share the mtime 0
	if err = a.Migrate(); err != nil {
		t.Fatal(err)
	}

	count, err := a.CountBlocks()
	if err != nil {
		t.Fatal(err)
	}
	if count < 20 {
		t.Fatalf("not enough blocks in the test database: %d", count)
	}

	// every page sorts the blocks with the same mtime: keep the page count low
	for _, limit := range []int{count/20 + 1, count/3 + 1, count - 1, count, count + 1} {
		list := findAllBlocksByMtime(t, a, limit)
		if len(list) != count {
			t.Errorf("limit %d: got %d of %d blocks", limit, len(list), count)
		}
	}

	// touch a few blocks, the update trigger sets their mtime to the current time: they have to come last
	rwdb, err := sql.Open("sqlite", tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer rwdb.Close()

	all := findAllBlocksByMtime(t, a, count)
	newer := map[types.MapBlockCoords]bool{}
	for _, pos := range all[:5] {
		newer[pos] = true
		if a.legacy_pos {
			_, err = rwdb.Exec("update blocks set mtime = 1 where pos = ?", util.CoordToPlain(&pos))
		} else {
			_, err = rwdb.Exec("update blocks set mtime = 1 where x = ? and y = ? and z = ?", pos.X, pos.Y, pos.Z)
		}
		if err != nil {
			t.Fatal(err)
		}
	}

	list := findAllBlocksByMtime(t, a, count/20+1)
	if len(list) != count {
		t.Fatalf("got %d of %d blocks", len(list), count)
	}
	for i, pos := range list {
		if newer[pos] != (i >= count-5) {
			t.Errorf("block %v at unexpected index %d", pos, i)
		}
	}

	// without a position only blocks with a newer mtime are returned
	blocks, err := a.FindBlocksByMtime(0, nil, count)
	if err != nil {
		t.Fatal(err)
	}
	if len(blocks) != 5 {
		t.Errorf("got %d blocks newer than mtime 0, expected 5", len(blocks))
	}

	var newest int64
	for _, block := range blocks {
		if block.Mtime > newest {
			newest = block.Mtime
		}
	}

	blocks, err = a.FindBlocksByMtime(newest, nil, count)
	if err != nil {
		t.Fatal(err)
	}
	if len(blocks) != 0 {
		t.Errorf("got %d blocks newer than the newest mtime, expected none", len(blocks))
	}
}

func TestFindBlocksByMtimePagination(t *testing.T) {
	testFindBlocksByMtimePagination(t, testutils.CreateTestDatabase)
}

func TestFindBlocksByMtimePaginationLegacy(t *testing.T) {
	testFindBlocksByMtimePagination(t, testutils.CreateTestDatabaseLegacy)
}
