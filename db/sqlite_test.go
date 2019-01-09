package db

import (
  "os"
  "io"
  "fmt"
  "testing"
  "io/ioutil"
  "database/sql"
  "mapserver/coords"
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
  tmpfile, err := ioutil.TempFile("", "TestMigrateEmpty.*.sqlite")
  if err != nil {
    panic(err)
  }
  defer os.Remove(tmpfile.Name())

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
  tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
  if err != nil {
    panic(err)
  }
  defer os.Remove(tmpfile.Name())

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


func TestMigrateAndQuery(t *testing.T){
  tmpfile, err := ioutil.TempFile("", "TestMigrateAndQuery.*.sqlite")
  if err != nil {
    panic(err)
  }
  defer os.Remove(tmpfile.Name())

  createTestDatabase(tmpfile.Name())
  a, err := NewSqliteAccessor(tmpfile.Name())
  if err != nil {
    panic(err)
  }
  err = a.Migrate()
  if err != nil {
    panic(err)
  }

  block, err := a.GetBlock(coords.NewMapBlockCoords(0,0,0))

  if err != nil {
    panic(err)
  }

  if block == nil {
    t.Fatal("no data")
  }

}
