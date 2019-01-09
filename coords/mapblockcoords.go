package coords

type MapBlockCoords struct {
  X,Y,Z int
}

func NewMapBlockCoords(x,y,z int) MapBlockCoords {
  return MapBlockCoords{X:x, Y:y, Z:z}
}
