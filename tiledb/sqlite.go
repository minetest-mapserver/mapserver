package tiledb

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/sirupsen/logrus"
  "mapserver/coords"
  "time"
)

const migrateScript = `
create table if not exists tiles(
  data blob,
  mtime bigint,
  layer int,
  x int,
  y int,
  zoom int,
  primary key(x,y,zoom,layer)
);
`

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {
		log.WithFields(logrus.Fields{"filename": db.filename}).Info("Migrating database")
		start := time.Now()
		_, err := db.db.Exec(migrateScript)
		if err != nil {
			return err
		}
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")

  	return nil
}

func (db *Sqlite3Accessor) GetTile(pos coords.TileCoords) (*Tile, error) {
  return nil, nil
}

func (db *Sqlite3Accessor) SetTile(pos coords.TileCoords, tile *Tile) error {
  return nil
}

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
