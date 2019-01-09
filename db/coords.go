package db

//https://bitbucket.org/s_l_teichmann/mtsatellite/src/e1bf980a2b278c570b3f44f9452c9c087558acb3/common/coords.go?at=default&fileviewer=file-view-default
const (
	numBitsPerComponent = 12
	modulo              = 1 << numBitsPerComponent
	maxPositive         = modulo / 2
	minValue            = -1 << (numBitsPerComponent - 1)
	maxValue            = 1<<(numBitsPerComponent-1) - 1
)


func CoordToPlain(x, y, z int) int64 {
	return int64(z)<<(2*numBitsPerComponent) +
		int64(y)<<numBitsPerComponent +
		int64(x)
}

func unsignedToSigned(i int16) int16 {
	if i < maxPositive {
		return i
	}
	return i - maxPositive*2
}

// To match C++ code.
func pythonModulo(i int16) int16 {
	const mask = modulo - 1
	if i >= 0 {
		return i & mask
	}
	return modulo - -i&mask
}

func PlainToCoord(i int64) (int, int, int) {
	x := unsignedToSigned(pythonModulo(int16(i)))
	i = (i - int64(x)) >> numBitsPerComponent
	y := unsignedToSigned(pythonModulo(int16(i)))
	i = (i - int64(x)) >> numBitsPerComponent
	z := unsignedToSigned(pythonModulo(int16(i)))
	return int(x), int(y), int(z)
}
