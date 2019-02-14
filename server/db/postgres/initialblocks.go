package postgres

import (
	"github.com/sirupsen/logrus"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/layer"
	"mapserver/settings"
)

const (
	SETTING_LAST_LAYER   = "last_layer"
	SETTING_LAST_X_BLOCK = "last_x_block"
	SETTING_LAST_Y_BLOCK = "last_y_block"
)

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
	lastxblock := s.GetInt(SETTING_LAST_X_BLOCK, -128)
	lastyblock := s.GetInt(SETTING_LAST_Y_BLOCK, -128)

	if lastxblock >= 128 {
		lastxblock = -128
		lastyblock++

	} else {
		lastxblock++

	}

	if lastyblock > 128 {
		//done
		//TODO: next layer

		result := &db.InitialBlocksResult{}
		result.HasMore = false
		return result, nil
	}

	tc := coords.NewTileCoords(lastxblock, lastyblock, 9, lastlayer)
	currentlayer := layer.FindLayerById(layers, lastlayer)

	tcr := coords.GetMapBlockRangeFromTile(tc, currentlayer.From)

	fields := logrus.Fields{
		"layerId": lastlayer,
		"pos1":    tcr.Pos1,
		"pos2":    tcr.Pos2,
	}
	log.WithFields(fields).Info("Initial-Query")

	rows, err := this.db.Query(getBlocksByInitialTileQuery,
		tcr.Pos1.X, tcr.Pos1.X, tcr.Pos1.X,
		tcr.Pos2.X, tcr.Pos2.X, tcr.Pos2.X,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	blocks := make([]*db.Block, 0)

	for rows.Next() {
		var posx, posy, posz int
		var data []byte
		var mtime int64

		err = rows.Scan(&posx, &posy, &posz, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(posx, posy, posz, data, mtime)
		blocks = append(blocks, mb)
	}

	result := &db.InitialBlocksResult{}
	result.List = blocks
	result.HasMore = true

	return result, nil
}
