package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetZoomedOutTile(t *testing.T) {
	tests := []struct {
		in, out TileCoords
	}{
		{TileCoords{X: 0, Y: 0, Zoom: 13, LayerId: 1}, TileCoords{X: 0, Y: 0, Zoom: 12, LayerId: 1}},
		{TileCoords{X: 5, Y: 7, Zoom: 13, LayerId: 2}, TileCoords{X: 2, Y: 3, Zoom: 12, LayerId: 2}},
		// negative coordinates floor towards -inf
		{TileCoords{X: -1, Y: -1, Zoom: 13}, TileCoords{X: -1, Y: -1, Zoom: 12}},
		{TileCoords{X: -2, Y: -3, Zoom: 13}, TileCoords{X: -1, Y: -2, Zoom: 12}},
	}

	for _, tc := range tests {
		assert.Equal(t, &tc.out, tc.in.GetZoomedOutTile())
	}
}

func TestZoomOut(t *testing.T) {
	tc := NewTileCoords(13, -9, 13, 3)

	assert.Equal(t, tc, tc.ZoomOut(0))
	assert.Equal(t, NewTileCoords(3, -3, 11, 3), tc.ZoomOut(2))
	// original is not modified
	assert.Equal(t, NewTileCoords(13, -9, 13, 3), tc)
}

func TestGetZoomedQuadrantsFromTile(t *testing.T) {
	q := NewTileCoords(-1, 2, 5, 4).GetZoomedQuadrantsFromTile()

	assert.Equal(t, NewTileCoords(-2, 4, 6, 4), q.UpperLeft)
	assert.Equal(t, NewTileCoords(-1, 4, 6, 4), q.UpperRight)
	assert.Equal(t, NewTileCoords(-2, 5, 6, 4), q.LowerLeft)
	assert.Equal(t, NewTileCoords(-1, 5, 6, 4), q.LowerRight)
}

func TestQuadrantsZoomOutRoundtrip(t *testing.T) {
	parent := NewTileCoords(-3, 8, 10, 1)
	q := parent.GetZoomedQuadrantsFromTile()

	for _, child := range []*TileCoords{q.UpperLeft, q.UpperRight, q.LowerLeft, q.LowerRight} {
		assert.Equal(t, parent, child.GetZoomedOutTile())
	}
}
