package sqlite

import (
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layer"
	"mapserver/settings"
)

const (
	SETTING_LAST_POS               = "last_pos"
	SETTING_TOTAL_LEGACY_COUNT     = "total_legacy_count"
	SETTING_PROCESSED_LEGACY_COUNT = "total_processed_legacy_count"
)

const getLastBlockQuery = `
select pos,data,mtime
from blocks b
where b.pos > ?
order by b.pos asc, b.mtime asc
limit ?
`

func (a *Sqlite3Accessor) FindNextInitialBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*db.InitialBlocksResult, error) {
	result := &db.InitialBlocksResult{}

	blocks := make([]*db.Block, 0)
	lastpos := s.GetInt64(SETTING_LAST_POS, coords.MinPlainCoord-1)

	processedcount := s.GetInt64(SETTING_PROCESSED_LEGACY_COUNT, 0)
	totallegacycount := s.GetInt64(SETTING_TOTAL_LEGACY_COUNT, -1)
	if totallegacycount == -1 {
		//Query from db
		totallegacycount, err := a.CountBlocks()

		if err != nil {
			panic(err)
		}

		s.SetInt64(SETTING_TOTAL_LEGACY_COUNT, int64(totallegacycount))
	}

	rows, err := a.db.Query(getLastBlockQuery, lastpos, limit)
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

		if mtime > result.LastMtime {
			result.LastMtime = mtime
		}

		mb := convertRows(pos, data, mtime)

		// new position
		lastpos = pos

		blockcoordy := mb.Pos.Y
		currentlayer := layer.FindLayerByY(layers, blockcoordy)

		if currentlayer != nil {
			blocks = append(blocks, mb)
		}
	}

	s.SetInt64(SETTING_PROCESSED_LEGACY_COUNT, int64(result.UnfilteredCount)+processedcount)

	result.Progress = float64(processedcount) / float64(totallegacycount)
	result.List = blocks

	//Save current positions of initial run
	s.SetInt64(SETTING_LAST_POS, lastpos)

	return result, nil
}
