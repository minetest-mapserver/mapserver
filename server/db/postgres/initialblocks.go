package postgres

import (
	"mapserver/db"
	"mapserver/layer"
	"mapserver/settings"
)

const (
	SETTING_LAST_LAYER = "last_layer"
)

func (this *PostgresAccessor) FindNextInitialBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*db.InitialBlocksResult, error) {

	//zoom/mapblock-width
	//13 1
	//12 2
	//11 4
	//10 8
	//9 16

	//Zoom 9:
	//10 mapblocks height * 16 * 16 == 2560
	return nil, nil
}
