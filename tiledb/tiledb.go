package tiledb

import (
	"errors"
	"fmt"
	"mapserver/util"
	"os"
)

func New(path string) (*TileDB, error) {
	return &TileDB{
		path: path,
	}, nil
}

type TileDB struct {
	path string
}

func (tdb *TileDB) getDirAndFile(pos *util.TileCoords) (string, string) {
	dir := fmt.Sprintf("%s/%d/%d/%d", tdb.path, pos.LayerId, pos.Zoom, pos.X)
	file := fmt.Sprintf("%s/%d.png", dir, pos.Y)
	return dir, file
}

func (tdb *TileDB) GetTile(pos *util.TileCoords) ([]byte, error) {
	_, file := tdb.getDirAndFile(pos)
	content, err := os.ReadFile(file)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (tdb *TileDB) SetTile(pos *util.TileCoords, tile []byte) error {
	dir, file := tdb.getDirAndFile(pos)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	return os.WriteFile(file, tile, 0644)
}
