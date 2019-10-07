package coords

import (
	"math"
)

const(
	MAX_ZOOM = 13
	MIN_ZOOM = 0
)

type TileCoords struct {
	X       int `json:"x"`
	Y       int `json:"y"`
	Zoom    int `json:"zoom"`
	LayerId int `json:"layerid"`
}

type TileQuadrants struct {
	UpperLeft, UpperRight, LowerLeft, LowerRight *TileCoords
}

func NewTileCoords(x, y, zoom int, layerId int) *TileCoords {
	return &TileCoords{X: x, Y: y, Zoom: zoom, LayerId: layerId}
}

func (tc *TileCoords) ZoomOut(n int) *TileCoords {
	var nc *TileCoords = tc
	for i := 0; i < n; i++ {
		nc = nc.GetZoomedOutTile()
	}

	return nc
}

func (tc *TileCoords) GetZoomedOutTile() *TileCoords {
	return &TileCoords{
		X:       int(math.Floor(float64(tc.X) / 2.0)),
		Y:       int(math.Floor(float64(tc.Y) / 2.0)),
		Zoom:    tc.Zoom - 1,
		LayerId: tc.LayerId}
}

func (tc *TileCoords) GetZoomedQuadrantsFromTile() TileQuadrants {
	nextZoom := tc.Zoom + 1

	nextZoomX := tc.X * 2
	nextZoomY := tc.Y * 2

	upperLeft := &TileCoords{X: nextZoomX, Y: nextZoomY, Zoom: nextZoom, LayerId: tc.LayerId}
	upperRight := &TileCoords{X: nextZoomX + 1, Y: nextZoomY, Zoom: nextZoom, LayerId: tc.LayerId}
	lowerLeft := &TileCoords{X: nextZoomX, Y: nextZoomY + 1, Zoom: nextZoom, LayerId: tc.LayerId}
	lowerRight := &TileCoords{X: nextZoomX + 1, Y: nextZoomY + 1, Zoom: nextZoom, LayerId: tc.LayerId}

	return TileQuadrants{
		UpperLeft:  upperLeft,
		UpperRight: upperRight,
		LowerLeft:  lowerLeft,
		LowerRight: lowerRight,
	}
}
