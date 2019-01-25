package playerdb

import (
	"errors"
	"mapserver/worldconfig"
)

func Create(cfg *worldconfig.WorldConfig) (DBAccessor, error) {
	switch cfg.PlayerBackend {
	case worldconfig.BACKEND_FILES:
		return &FilePlayerDB{}, nil
	default:
		return nil, errors.New("player backend not supported: " + worldconfig.BACKEND_FILES)
	}
}
