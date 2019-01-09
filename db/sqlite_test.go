package db

import (
  "os"
  "io"
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

const testDatabase = "./testdata/map.sqlite"

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

func createTestDatabase(filename string) error {
  return copy(testDatabase, filename)
}

func TestMigrateEmpty(t *testing.T){
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

func TestMigrate(t *testing.T){
  tmpfile, err := ioutil.TempFile("", "example")
  if err != nil {
    panic(err)
  }

  createTestDatabase(tmpfile.Name())
  a, err := NewSqliteAccessor(tmpfile.Name())
  if err != nil {
    panic(err)
  }
  err = a.Migrate()
  if err != nil {
    panic(err)
  }
}
