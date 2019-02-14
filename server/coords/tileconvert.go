package coords

import (
	"mapserver/layer"
	"math"
)

const (
	MAX_ZOOM = 13
)

func GetTileCoordsFromMapBlock(mbc *MapBlockCoords, layers []*layer.Layer) *TileCoords {
	tc := TileCoords{X: mbc.X, Y: (mbc.Z + 1) * -1, Zoom: MAX_ZOOM}

	var layerid *int
	for _, l := range layers {
		if (mbc.Y*16) >= l.From && (mbc.Y*16) <= l.To {
			layerid = &l.Id
			break
		}
	}

	if layerid == nil {
		return nil
	}

	tc.LayerId = *layerid

	return &tc
}

func GetMapBlockRangeFromTile(tc *TileCoords, y int) MapBlockRange {
	scaleDiff := float64(MAX_ZOOM - tc.Zoom)
	scale := int(math.Pow(2, scaleDiff))

	mapBlockX1 := tc.X * scale
	mapBlockZ1 := (tc.Y * scale * -1) - 1

	mapBlockX2 := mapBlockX1 + scale - 1
	mapBlockZ2 := (mapBlockZ1 + ((scale - 1) * -1))

	return MapBlockRange{
		Pos1: NewMapBlockCoords(mapBlockX1, y, mapBlockZ1),
		Pos2: NewMapBlockCoords(mapBlockX2, y, mapBlockZ2),
	}
}
