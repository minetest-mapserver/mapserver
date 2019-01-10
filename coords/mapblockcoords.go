package coords

type MapBlockCoords struct {
  X,Y,Z int
}

func NewMapBlockCoords(x,y,z int) MapBlockCoords {
  return MapBlockCoords{X:x, Y:y, Z:z}
}

type MapBlockRange struct {
  pos1, pos2 MapBlockCoords
}

const (
  MaxCoord = 2047
  MinCoord = -2047
)
