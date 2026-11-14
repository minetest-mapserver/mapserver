package types

import (
	"fmt"
	"math"
)

func GetChunkKey(x, y, z int) string {
	return fmt.Sprintf("%d/%d/%d", x, y, z)
}

func floorOfDivision(n, q int) int {
	return int(math.Floor(float64(n) / float64(q)))
}

func GetMapblockPosFromNode(x, y, z int) (int, int, int) {
	return floorOfDivision(x, 16), floorOfDivision(y, 16), floorOfDivision(z, 16)
}

func GetMapblockBoundsFromChunk(x, y, z int) (x1, y1, z1, x2, y2, z2 int) {
	x1 = (x * 5)
	y1 = (y * 5)
	z1 = (z * 5)
	x2 = x1 + 4
	y2 = y1 + 4
	z2 = z1 + 4
	return
}

func GetChunkPosFromMapblock(x, y, z int) (int, int, int) {
	return floorOfDivision((x), 5), floorOfDivision((y), 5), floorOfDivision((z), 5)
}

func GetChunkPosFromNode(x, y, z int) (int, int, int) {
	m_x, m_y, m_z := GetMapblockPosFromNode(x, y, z)
	return GetChunkPosFromMapblock(m_x, m_y, m_z)
}
