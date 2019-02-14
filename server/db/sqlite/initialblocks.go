package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/settings"
	"mapserver/layer"
)

const (
	SETTING_LAST_POS            = "last_pos"
)

const getLastBlockQuery = `
select pos,data,mtime
from blocks b
where b.mtime = 0
and b.pos > ?
order by b.pos asc, b.mtime asc
limit ?
`

func (this *Sqlite3Accessor) FindNextInitialBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*db.InitialBlocksResult, error) {

	result := &db.InitialBlocksResult{}

	blocks := make([]*db.Block, 0)

	lastpos := s.GetInt64(SETTING_LAST_POS, coords.MinPlainCoord-1)

	rows, err := this.db.Query(getLastBlockQuery, lastpos, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		result.HasMore = true
		result.UnfilteredCount++

		var pos int64
		var data []byte
		var mtime int64

		err = rows.Scan(&pos, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(pos, data, mtime)

		// new position
		lastpos = pos

		blockcoordy := mb.Pos.Y * 16
		currentlayer := layer.FindLayerByY(layers, blockcoordy)

		if currentlayer != nil {
			blocks = append(blocks, mb)
		}
	}

	result.List = blocks

	//Save current positions of initial run
	s.SetInt64(SETTING_LAST_POS, lastpos)

	return result, nil
}
