package coords

import (
	"mapserver/layer"
)

func GetTileCoordsFromMapBlock(mbc *MapBlockCoords, layers []*layer.Layer) *TileCoords {
	tc := TileCoords{X: mbc.X, Y: (mbc.Z + 1) * -1, Zoom: 0}

	currentLayer := layer.FindLayerByY(layers, mbc.Y)

	if currentLayer == nil {
		return nil
	}

	tc.LayerId = currentLayer.Id

	return &tc
}

func GetMapBlockRangeFromTile(tc *TileCoords, y int) *MapBlockRange {
	scale := 1
	if tc.Zoom > 0 {
		scale = 2 << (tc.Zoom - 1)
	}

	mapBlockX1 := tc.X * scale
	mapBlockZ1 := -(tc.Y * scale) - 1

	mapBlockX2 := mapBlockX1 + scale - 1
	mapBlockZ2 := (mapBlockZ1 - (scale - 1))

	return &MapBlockRange{
		Pos1: NewMapBlockCoords(mapBlockX1, y, mapBlockZ1),
		Pos2: NewMapBlockCoords(mapBlockX2, y, mapBlockZ2),
	}
}
