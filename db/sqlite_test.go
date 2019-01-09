package db

import (
  "fmt"
  "testing"
  "io/ioutil"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

const emptyBlocksScript = `
create table blocks (
  pos int,
  data blob
);
`

func createEmptyDatabase(filename string){
  db, err := sql.Open("sqlite3", filename)
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

func TestMigrate(t *testing.T){
  tmpfile, err := ioutil.TempFile("", "example")
  if err != nil {
    panic(err)
  }

  createEmptyDatabase(tmpfile.Name())
  a, err := NewSqliteAccessor(tmpfile.Name())
  if err != nil {
    panic(err)
  }
  err = a.Migrate()
  if err != nil {
    panic(err)
  }
}
