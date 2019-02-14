package postgres

import (
	"mapserver/settings"
	"mapserver/layer"
	"mapserver/db"
)



func (this *PostgresAccessor) FindNextInitialBlocks(s settings.Settings, layers []*layer.Layer, limit int) (*db.InitialBlocksResult, error) {
	return nil, nil
}
