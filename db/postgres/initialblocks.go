package postgres

import (
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layer"
	"mapserver/settings"
	"math"

	"github.com/sirupsen/logrus"
)

const (
	SETTING_LAST_LAYER   = "last_layer"
	SETTING_LAST_X_BLOCK = "last_x_block"
	SETTING_LAST_Y_BLOCK = "last_y_block"
)

func (this *PostgresAccessor) countBlocks(x1, y1, z1, x2, y2, z2 int) (int, error) {
	rows, err := this.db.Query(getBlockCountByInitialTileQuery,
		x1, y1, z1, x2, y2, z2,
	)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var count int

		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	return 0, nil
}

// x -> 0 ... 256

//zoom/mapblock-width
//13 1
//12 2
//11 4
//10 8
//9 16

//Zoom 9:
//10 mapblocks height * 16 * 16 == 2560

func (this *PostgresAccessor) FindNextInitialBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*db.InitialBlocksResult, error) {

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
	currentlayer := layer.FindLayerById(layers, lastlayer)

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
		stridecount := this.intQuery(`
			select count(*) from blocks
			where posz >= $1 and posz <= $2
			and posy >= $3 and posy <= $4`,
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

	count, err := this.countBlocks(minX, minY, minZ, maxX, maxY, maxZ)

	if err != nil {
		return nil, err
	}

	blocks := make([]*db.Block, 0)
	var lastmtime int64

	if count > 0 {

		rows, err := this.db.Query(getBlocksByInitialTileQuery,
			minX, minY, minZ, maxX, maxY, maxZ,
		)

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var posx, posy, posz int
			var data []byte
			var mtime int64

			err = rows.Scan(&posx, &posy, &posz, &data, &mtime)
			if err != nil {
				return nil, err
			}

			if mtime > lastmtime {
				lastmtime = mtime
			}

			mb := convertRows(posx, posy, posz, data, mtime)
			blocks = append(blocks, mb)
		}
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
