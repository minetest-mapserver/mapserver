package tiledb

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mapserver/coords"
	"os"
	"strconv"
	"sync"
)

var mutex = &sync.RWMutex{}

func New(path string) (*TileDB, error) {
	return &TileDB{
		path: path,
		outdated: make(map[int]map[int]map[int]map[int]bool),
	}, nil
}

type TileDB struct {
	path string
	outdated map[int]map[int]map[int]map[int]bool
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

func (this *TileDB) TileExists(pos *coords.TileCoords) (bool) {
	_, file := this.getDirAndFile(pos)
	info, _ := os.Stat(file)
	return info != nil
}

func (this *TileDB) GetTile(pos *coords.TileCoords) ([]byte, error) {
	timer := prometheus.NewTimer(tiledbLoadDuration)
	defer timer.ObserveDuration()

	mutex.RLock()
	defer mutex.RUnlock()

	fields := logrus.Fields{
		"pos": pos,
	}
	log.WithFields(fields).Debug("GetTile")

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
	timer := prometheus.NewTimer(tiledbSaveDuration)
	defer timer.ObserveDuration()

	mutex.Lock()
	defer mutex.Unlock()

	fields := logrus.Fields{
		"pos":  pos,
		"size": len(tile),
	}
	log.WithFields(fields).Debug("SetTile")

	dir, file := this.getDirAndFile(pos)
	os.MkdirAll(dir, 0700)

	err := ioutil.WriteFile(file, tile, 0644)

	if err == nil {
		this.setOutdated(pos, false)
	}

	return err
}

func (this *TileDB) MarkOutdated(pos *coords.TileCoords) {
	var npos *coords.TileCoords = pos
	mutex.RLock()
	defer mutex.RUnlock()

	for npos.Zoom <= 13 && !this.IsOutdated(npos) {
		this.setOutdated(npos, true)
		npos = npos.GetZoomedOutTile()
	}
}

func (this *TileDB) IsOutdated(pos *coords.TileCoords) bool {
	if this.outdated[pos.Zoom] == nil {
		return false
	}
	if this.outdated[pos.Zoom][pos.LayerId] == nil {
		return false
	}
	if this.outdated[pos.Zoom][pos.LayerId][pos.X] == nil {
		return false
	}
	return this.outdated[pos.Zoom][pos.LayerId][pos.X][pos.Y]
}

func (this *TileDB) setOutdated(pos *coords.TileCoords, outdated bool) {
	if this.outdated[pos.Zoom] == nil {
		this.outdated[pos.Zoom] = make(map[int]map[int]map[int]bool)
	}
	if this.outdated[pos.Zoom][pos.LayerId] == nil {
		this.outdated[pos.Zoom][pos.LayerId] = make(map[int]map[int]bool)
	}
	if this.outdated[pos.Zoom][pos.LayerId][pos.X] == nil {
		this.outdated[pos.Zoom][pos.LayerId][pos.X] = make(map[int]bool)
	}
	this.outdated[pos.Zoom][pos.LayerId][pos.X][pos.Y] = outdated
}

func (this *TileDB) GetOutdatedByZoom(zoom int) []coords.TileCoords {
	var tiles []coords.TileCoords
	if this.outdated[zoom] == nil {
		return tiles
	}
	for layerId, layerContent := range this.outdated[zoom] {
		for x, xContent := range layerContent {
			for y, outdated := range xContent {
				if outdated {
					tiles = append(tiles,
							coords.TileCoords{X:x, Y:y, Zoom:zoom, LayerId:layerId})
				}
			}
		}
	}
	return tiles
}
