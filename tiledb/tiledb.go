package tiledb

import (
	"fmt"
	"mapserver/coords"
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

func (tdb *TileDB) getDirAndFile(pos *coords.TileCoords) (string, string) {
	dir := fmt.Sprintf("%s/%d/%d/%d", tdb.path, pos.LayerId, pos.Zoom, pos.X)
	file := fmt.Sprintf("%s/%d.png", dir, pos.Y)
	return dir, file
}

func (tdb *TileDB) GetTile(pos *coords.TileCoords) ([]byte, error) {
	_, file := tdb.getDirAndFile(pos)
	info, _ := os.Stat(file)
	if info != nil {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		return content, err
	}

	return nil, nil
}

func (tdb *TileDB) SetTile(pos *coords.TileCoords, tile []byte) error {
	dir, file := tdb.getDirAndFile(pos)
	os.MkdirAll(dir, 0700)

	err := os.WriteFile(file, tile, 0644)
	return err
}
