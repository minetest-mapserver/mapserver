package coords

import (
	"math"
)

type TileCoords struct {
	X, Y int
	Zoom int
}

type TileQuadrants struct {
	UpperLeft, UpperRight, LowerLeft, LowerRight TileCoords
}

func NewTileCoords(x, y, zoom int) TileCoords {
	return TileCoords{X: x, Y: y, Zoom: zoom}
}

func (tc TileCoords) GetZoomedOutTile() TileCoords {
	return TileCoords{
		X:    int(math.Floor(float64(tc.X) / 2.0)),
		Y:    int(math.Floor(float64(tc.Y) / 2.0)),
		Zoom: tc.Zoom - 1}
}

func (tc TileCoords) GetZoomedQuadrantsFromTile() TileQuadrants {
	nextZoom := tc.Zoom + 1

	nextZoomX := tc.X * 2
	nextZoomY := tc.Y * 2

	upperLeft := TileCoords{X: nextZoomX, Y: nextZoomY, Zoom: nextZoom}
	upperRight := TileCoords{X: nextZoomX + 1, Y: nextZoomY, Zoom: nextZoom}
	lowerLeft := TileCoords{X: nextZoomX, Y: nextZoomY + 1, Zoom: nextZoom}
	lowerRight := TileCoords{X: nextZoomX + 1, Y: nextZoomY + 1, Zoom: nextZoom}

	return TileQuadrants{
		UpperLeft:  upperLeft,
		UpperRight: upperRight,
		LowerLeft:  lowerLeft,
		LowerRight: lowerRight,
	}
}
