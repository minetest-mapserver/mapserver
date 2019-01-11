package colormapping

import (
  "mapserver/vfs"
)

type ColorMapping struct {

}

func (m *ColorMapping) LoadColors(filename string){
  //TODO
}

func CreateColorMapping() {
  //embedded colors
  vfs.FSMustByte(false, "/colors.txt")
}
