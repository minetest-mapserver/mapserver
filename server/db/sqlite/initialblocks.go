package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"mapserver/coords"
	"mapserver/db"
)

func (this *Sqlite3Accessor) FindNextInitialBlocks(lastpos *coords.MapBlockCoords, limit int) (*db.InitialBlocksResult, error) {

	result := &db.InitialBlocksResult{}

	blocks := make([]*db.Block, 0)
	pc := coords.CoordToPlain(lastpos)

	rows, err := this.db.Query(getLastBlockQuery, pc, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var pos int64
		var data []byte
		var mtime int64

		err = rows.Scan(&pos, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(pos, data, mtime)
		blocks = append(blocks, mb)
	}

	result.List = blocks

	return result, nil
}
