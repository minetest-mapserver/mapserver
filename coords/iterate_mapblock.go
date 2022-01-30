package coords

type MapblockIterator func(x, y, z int)

func IterateMapblock(it MapblockIterator) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				it(x, y, z)
			}
		}
	}
}
