package coords

import (
	"github.com/stretchr/testify/assert"
	"mapserver/layer"
	"testing"
)

func TestConvertMapblockToTile1(t *testing.T) {
	mbc := NewMapBlockCoords(0, 0, 0)
	layers := []layer.Layer{
		layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	tc := GetTileCoordsFromMapBlock(mbc, layers)

	if tc.X != 0 {
		t.Fatal("x does not match")
	}

	if tc.Y != -1 {
		t.Fatal("y does not match")
	}

	if tc.Zoom != 13 {
		t.Fatal("zoom does not match")
	}
}

func TestGetMapBlockRangeFromTile(t *testing.T) {
	r := GetMapBlockRangeFromTile(NewTileCoords(0, 0, 13, 0), 0)
	assert.Equal(t, r.Pos1.X, 0)
	assert.Equal(t, r.Pos1.Z, -1)
	assert.Equal(t, r.Pos2.X, 0)
	assert.Equal(t, r.Pos2.Z, -1)

	r = GetMapBlockRangeFromTile(NewTileCoords(-1, -1, 13, 0), 0)
	assert.Equal(t, r.Pos1.X, -1)
	assert.Equal(t, r.Pos1.Z, 0)
	assert.Equal(t, r.Pos2.X, -1)
	assert.Equal(t, r.Pos2.Z, 0)
}

func TestConvertMapblockToTile2(t *testing.T) {
	mbc := NewMapBlockCoords(1, 0, 1)
	layers := []layer.Layer{
		layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	tc := GetTileCoordsFromMapBlock(mbc, layers)

	if tc.X != 1 {
		t.Fatal("x does not match")
	}

	if tc.Y != -2 {
		t.Fatal("y does not match")
	}

	if tc.Zoom != 13 {
		t.Fatal("zoom does not match")
	}
}

func TestConvertMapblockToTile3(t *testing.T) {
	mbc := NewMapBlockCoords(-1, 0, -1)
	layers := []layer.Layer{
		layer.Layer{
			Id:   0,
			Name: "Base",
			From: -16,
			To:   160,
		},
	}

	tc := GetTileCoordsFromMapBlock(mbc, layers)

	if tc.X != -1 {
		t.Fatal("x does not match")
	}

	if tc.Y != 0 {
		t.Fatal("y does not match")
	}

	if tc.Zoom != 13 {
		t.Fatal("zoom does not match")
	}
}

func TestZoomedQuadrantsFromTile(t *testing.T) {
	tc := NewTileCoords(0, 0, 12, 0)
	q := tc.GetZoomedQuadrantsFromTile()

	assert.Equal(t, q.UpperLeft.X, 0)
	assert.Equal(t, q.UpperLeft.Y, 0)
	assert.Equal(t, q.UpperLeft.Zoom, 13)

	assert.Equal(t, q.UpperRight.X, 1)
	assert.Equal(t, q.UpperRight.Y, 0)
	assert.Equal(t, q.UpperRight.Zoom, 13)

	assert.Equal(t, q.LowerLeft.X, 0)
	assert.Equal(t, q.LowerLeft.Y, 1)
	assert.Equal(t, q.LowerLeft.Zoom, 13)

	assert.Equal(t, q.LowerRight.X, 1)
	assert.Equal(t, q.LowerRight.Y, 1)
	assert.Equal(t, q.LowerRight.Zoom, 13)

}
