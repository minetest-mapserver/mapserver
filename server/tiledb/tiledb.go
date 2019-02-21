package tiledb

import (
	"io/ioutil"
	"mapserver/coords"
	"os"
	"strconv"
)

func New(path string) (*TileDB, error) {
	return &TileDB{
		path: path,
	}, nil
}

type TileDB struct {
	path string
}

func (this *TileDB) getDirAndFile(pos *coords.TileCoords) (string, string) {
	dir := this.path + "/" +
		strconv.Itoa(pos.LayerId) + "/" +
		strconv.Itoa(pos.Zoom) + "/" +
		strconv.Itoa(pos.X)

	file := dir + "/" + strconv.Itoa(pos.Y) + ".png"
	return dir, file
}

func (this *TileDB) GC() {
	//No-Op
}

func (this *TileDB) GetTile(pos *coords.TileCoords) ([]byte, error) {
	_, file := this.getDirAndFile(pos)
	info, _ := os.Stat(file)
	if info != nil {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		return content, err
	}

	return nil, nil
}

func (this *TileDB) SetTile(pos *coords.TileCoords, tile []byte) error {
	dir, file := this.getDirAndFile(pos)
	os.MkdirAll(dir, 0700)

	err := ioutil.WriteFile(file, tile, 0644)

	return err
}
