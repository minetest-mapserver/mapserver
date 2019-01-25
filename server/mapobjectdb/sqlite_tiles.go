package mapobjectdb

import (
	"mapserver/coords"
)

const getTileQuery = `
select data,mtime from tiles t
where t.layerid = ?
and t.x = ?
and t.y = ?
and t.zoom = ?
`

func (db *Sqlite3Accessor) GetTile(pos *coords.TileCoords) (*Tile, error) {
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
			Pos:   pos,
			Data:  data,
			Mtime: mtime,
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

func (db *Sqlite3Accessor) RemoveTile(pos *coords.TileCoords) error {
	_, err := db.db.Exec(removeTileQuery, pos.X, pos.Y, pos.Zoom, pos.LayerId)
	return err
}
