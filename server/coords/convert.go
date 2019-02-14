package coords

//https://bitbucket.org/s_l_teichmann/mtsatellite/src/e1bf980a2b278c570b3f44f9452c9c087558acb3/common/coords.go?at=default&fileviewer=file-view-default
const (
	numBitsPerComponent = 12
	modulo              = 1 << numBitsPerComponent
	maxPositive         = modulo / 2
	minValue            = -1 << (numBitsPerComponent - 1)
	maxValue            = 1<<(numBitsPerComponent-1) - 1

	MinPlainCoord	= -34351347711
)

func CoordToPlain(c *MapBlockCoords) int64 {
	return int64(c.Z)<<(2*numBitsPerComponent) +
		int64(c.Y)<<numBitsPerComponent +
		int64(c.X)
}

func unsignedToSigned(i int16) int {
	if i < maxPositive {
		return int(i)
	}
	return int(i - maxPositive*2)
}

// To match C++ code.
func pythonModulo(i int16) int16 {
	const mask = modulo - 1
	if i >= 0 {
		return i & mask
	}
	return modulo - -i&mask
}

func PlainToCoord(i int64) *MapBlockCoords {
	c := MapBlockCoords{}
	c.X = unsignedToSigned(pythonModulo(int16(i)))
	i = (i - int64(c.X)) >> numBitsPerComponent
	c.Y = unsignedToSigned(pythonModulo(int16(i)))
	i = (i - int64(c.Y)) >> numBitsPerComponent
	c.Z = unsignedToSigned(pythonModulo(int16(i)))
	return &c
}
