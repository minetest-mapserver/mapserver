package mapblockparser


func IterateMapblock(cb func(x,y,z int)){
  for x := 0; x < 16; x++ {
    for y := 0; y < 16; y++ {
      for z := 0; z < 16; z++ {
        cb(x,y,z)
      }
    }
  }
}
