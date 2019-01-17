package coords

import (
	"math"
)

const (
	MAX_ZOOM = 13
)

func GetTileCoordsFromMapBlock(mbc MapBlockCoords) TileCoords {
	return TileCoords{X: mbc.X, Y: (mbc.Z + 1) * -1, Zoom: MAX_ZOOM}
}

func GetMapBlockRangeFromTile(tc TileCoords, y int) MapBlockRange {
	scaleDiff := float64(MAX_ZOOM - tc.Zoom)
	scale := int(math.Pow(2, scaleDiff))

	mapBlockX1 := tc.X * scale
	mapBlockZ1 := (tc.Y * scale * -1) - 1

	mapBlockX2 := mapBlockX1 + scale-1
	mapBlockZ2 := (mapBlockZ1 + ((scale-1) * -1))

	return MapBlockRange{
		Pos1: NewMapBlockCoords(mapBlockX1, y, mapBlockZ1),
		Pos2: NewMapBlockCoords(mapBlockX2, y, mapBlockZ2),
	}
}
