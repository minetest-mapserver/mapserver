package sqlite

import (
	"mapserver/coords"
	"mapserver/db"
	"mapserver/settings"
	"mapserver/types"
	"math"

	"github.com/sirupsen/logrus"
)

const (
	SETTING_LAST_POS               = "last_pos"
	SETTING_TOTAL_LEGACY_COUNT     = "total_legacy_count"
	SETTING_PROCESSED_LEGACY_COUNT = "total_processed_legacy_count"

	SETTING_LAST_LAYER   = "last_layer"
	SETTING_LAST_X_BLOCK = "last_x_block"
	SETTING_LAST_Y_BLOCK = "last_y_block"
)

func (a *Sqlite3Accessor) FindNextInitialBlocks(s settings.Settings, layers []*types.Layer, limit int) (*db.InitialBlocksResult, error) {
	if a.legacy_pos {
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

		rows, err := a.db.Query(getLastBlockQueryLegacy, lastpos, limit)
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
			currentlayer := types.FindLayerByY(layers, blockcoordy)

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
	} else {

		lastlayer := s.GetInt(SETTING_LAST_LAYER, 0)
		lastxblock := s.GetInt(SETTING_LAST_X_BLOCK, -129)
		lastyblock := s.GetInt(SETTING_LAST_Y_BLOCK, -128)

		if lastxblock >= 128 {
			lastxblock = -128
			lastyblock++

		} else {
			lastxblock++

		}

		if lastyblock > 128 {
			//done
			var nextlayer = lastlayer
			for _, l := range layers {
				if l.Id > nextlayer {
					nextlayer = l.Id
					break
				}
			}

			s.SetInt(SETTING_LAST_LAYER, nextlayer)
			s.SetInt(SETTING_LAST_X_BLOCK, -129)
			s.SetInt(SETTING_LAST_Y_BLOCK, -128)

			result := &db.InitialBlocksResult{}
			result.HasMore = nextlayer != lastlayer
			return result, nil
		}

		tc := coords.NewTileCoords(lastxblock, lastyblock, 9, lastlayer)
		currentlayer := types.FindLayerById(layers, lastlayer)

		tcr := coords.GetMapBlockRangeFromTile(tc, 0)
		tcr.Pos1.Y = currentlayer.From
		tcr.Pos2.Y = currentlayer.To

		fields := logrus.Fields{
			"layerId": lastlayer,
			"pos1":    tcr.Pos1,
			"pos2":    tcr.Pos2,
			"tile":    tc,
		}
		log.WithFields(fields).Debug("Initial-Query")

		minX := int(math.Min(float64(tcr.Pos1.X), float64(tcr.Pos2.X)))
		maxX := int(math.Max(float64(tcr.Pos1.X), float64(tcr.Pos2.X)))
		minY := int(math.Min(float64(tcr.Pos1.Y), float64(tcr.Pos2.Y)))
		maxY := int(math.Max(float64(tcr.Pos1.Y), float64(tcr.Pos2.Y)))
		minZ := int(math.Min(float64(tcr.Pos1.Z), float64(tcr.Pos2.Z)))
		maxZ := int(math.Max(float64(tcr.Pos1.Z), float64(tcr.Pos2.Z)))

		if lastxblock <= -128 {
			//first x entry, check z stride
			stridecount := a.intQuery(`
				select count(*) from blocks
				where z >= ? and z <= ?
				and y >= ? and y <= ?`,
				minZ, maxZ,
				minY, maxY,
			)

			if stridecount == 0 {
				fields = logrus.Fields{
					"minX": minX,
					"maxX": maxX,
					"minY": minY,
					"maxY": maxY,
				}
				log.WithFields(fields).Debug("Skipping stride")

				s.SetInt(SETTING_LAST_LAYER, lastlayer)
				s.SetInt(SETTING_LAST_X_BLOCK, -129)
				s.SetInt(SETTING_LAST_Y_BLOCK, lastyblock+1)

				result := &db.InitialBlocksResult{}
				result.Progress = float64(((lastyblock+128)*256)+(lastxblock+128)) / float64(256*256)
				result.HasMore = true
				return result, nil
			}
		}

		blocks := make([]*db.Block, 0)
		var lastmtime int64

		rows, err := a.db.Query(getBlocksByInitialTileQuery,
			minX, minY, minZ, maxX, maxY, maxZ,
		)

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			mb := &db.Block{Pos: &types.MapBlockCoords{}}
			err = rows.Scan(&mb.Pos.X, &mb.Pos.Y, &mb.Pos.Z, &mb.Data, &mb.Mtime)
			if err != nil {
				return nil, err
			}

			if mb.Mtime > lastmtime {
				lastmtime = mb.Mtime
			}

			blocks = append(blocks, mb)
		}

		s.SetInt(SETTING_LAST_LAYER, lastlayer)
		s.SetInt(SETTING_LAST_X_BLOCK, lastxblock)
		s.SetInt(SETTING_LAST_Y_BLOCK, lastyblock)

		result := &db.InitialBlocksResult{}
		result.LastMtime = lastmtime
		result.Progress = float64(((lastyblock+128)*256)+(lastxblock+128)) / float64(256*256)
		result.List = blocks
		result.HasMore = true

		return result, nil
	}
}
