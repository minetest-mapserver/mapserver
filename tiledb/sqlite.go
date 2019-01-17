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
  layerid int,
  x int,
  y int,
  zoom int,
  primary key(x,y,zoom,layerid)
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

const getTileQuery = `
select data,mtime from tiles t
where t.layerid = ?
and t.x = ?
and t.y = ?
and t.zoom = ?
`

func (db *Sqlite3Accessor) GetTile(pos coords.TileCoords) (*Tile, error) {
	rows, err := db.db.Query(getTileQuery, pos.LayerId, pos.X, pos.Y, pos.Zoom)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var data []byte
		var mtime int64

		err = rows.Scan(&data, &mtime)
		if err != nil {
			return nil, err
		}

		if data == nil {
			return nil, nil
		}

		mb := Tile{
			Pos:     pos,
			Data:    data,
			Mtime:   mtime,
		}

		return &mb, nil
	}

	return nil, nil
}

const setTileQuery = `
insert or replace into tiles(x,y,zoom,layerid,data,mtime)
values(?, ?, ?, ?, ?, ?)
`

func (db *Sqlite3Accessor) SetTile(tile *Tile) error {
	_, err := db.db.Exec(setTileQuery, tile.Pos.X, tile.Pos.Y, tile.Pos.Zoom, tile.Pos.LayerId, tile.Data, tile.Mtime)
	return err
}

const removeTileQuery = `
delete from tiles
where x = ? and y = ? and zoom = ? and layerid = ?
`

func (db *Sqlite3Accessor) RemoveTile(pos coords.TileCoords) error {
	_, err := db.db.Exec(removeTileQuery, pos.X, pos.Y, pos.Zoom, pos.LayerId)
	return err
}


func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
