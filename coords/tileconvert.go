package coords

import (
	"mapserver/types"
	"math"
)

const (
	MAX_ZOOM = 13
)

func GetTileCoordsFromMapBlock(mbc *types.MapBlockCoords, layers []*types.Layer) *TileCoords {
	tc := TileCoords{X: mbc.X, Y: (mbc.Z + 1) * -1, Zoom: MAX_ZOOM}

	currentLayer := types.FindLayerByY(layers, mbc.Y)

	if currentLayer == nil {
		return nil
	}

	tc.LayerId = currentLayer.Id

	return &tc
}

func GetMapBlockRangeFromTile(tc *TileCoords, y int) *types.MapBlockRange {
	scaleDiff := float64(MAX_ZOOM - tc.Zoom)
	scale := int(math.Pow(2, scaleDiff))

	mapBlockX1 := tc.X * scale
	mapBlockZ1 := (tc.Y * scale * -1) - 1

	mapBlockX2 := mapBlockX1 + scale - 1
	mapBlockZ2 := (mapBlockZ1 + ((scale - 1) * -1))

	return &types.MapBlockRange{
		Pos1: types.NewMapBlockCoords(mapBlockX1, y, mapBlockZ1),
		Pos2: types.NewMapBlockCoords(mapBlockX2, y, mapBlockZ2),
	}
}
