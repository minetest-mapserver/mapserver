package coords

import (
	"math"
)

type MapBlockCoords struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func NewMapBlockCoords(x, y, z int) *MapBlockCoords {
	return &MapBlockCoords{X: x, Y: y, Z: z}
}

func NewMapBlockCoordsFromBlock(x, y, z int) *MapBlockCoords {
	return &MapBlockCoords{
		X: int(math.Floor(float64(x) / 16)),
		Y: int(math.Floor(float64(y) / 16)),
		Z: int(math.Floor(float64(z) / 16)),
	}
}

type MapBlockRange struct {
	Pos1, Pos2 *MapBlockCoords
}

const (
	MaxCoord = 2047
	MinCoord = -2047
)
