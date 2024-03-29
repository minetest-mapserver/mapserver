package testutils

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	_ "modernc.org/sqlite"
)

const emptyBlocksScript = `
create table blocks (
  pos int,
  data blob
);
`

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func CreateTestDatabase(filename string) error {
	_, currentfilename, _, _ := runtime.Caller(0)
	return copy(filepath.Dir(currentfilename)+"/testdata/map1.sqlite", filename)
}

//DB with metadata at 0,0,0
func CreateTestDatabase2(filename string) error {
	_, currentfilename, _, _ := runtime.Caller(0)
	return copy(filepath.Dir(currentfilename)+"/testdata/map2.sqlite", filename)
}

func CreateEmptyDatabase(filename string) {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(emptyBlocksScript)
	if err != nil {
		panic(err)
	}
	rows.Next()
	fmt.Println(rows)
	db.Close()
}
