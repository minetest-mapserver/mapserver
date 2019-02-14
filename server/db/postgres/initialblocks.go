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
	return nil, nil
}
