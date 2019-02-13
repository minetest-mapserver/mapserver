package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/settings"
)

const getLastBlockQuery = `
select pos,data,mtime
from blocks b
where b.mtime = 0
and b.pos > ?
order by b.pos asc, b.mtime asc
limit ?
`

func (this *Sqlite3Accessor) FindNextInitialBlocks(s settings.Settings, layers []layer.Layer, limit int) (*db.InitialBlocksResult, error) {

	result := &db.InitialBlocksResult{}

	blocks := make([]*db.Block, 0)

	lastx := s.GetInt(settings.SETTING_LASTX, coords.MinCoord-1)
	lasty := s.GetInt(settings.SETTING_LASTY, coords.MinCoord-1)
	lastz := s.GetInt(settings.SETTING_LASTZ, coords.MinCoord-1)
	
	lastcoords := coords.NewMapBlockCoords(lastx, lasty, lastz)
	pc := coords.CoordToPlain(lastpos)

	rows, err := this.db.Query(getLastBlockQuery, pc, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var newlastpos *coords.MapBlockCoords

	for rows.Next() {
		var pos int64
		var data []byte
		var mtime int64

		err = rows.Scan(&pos, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(pos, data, mtime)
		newlastpos = mb.Pos

		blocks = append(blocks, mb)
	}

	result.List = blocks

	//Save current positions of initial run
	s.SetInt(settings.SETTING_LASTX, newlastpos.X)
	s.SetInt(settings.SETTING_LASTY, newlastpos.Y)
	s.SetInt(settings.SETTING_LASTZ, newlastpos.Z)

	return result, nil
}
