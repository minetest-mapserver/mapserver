package mapblockparser

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"strconv"
)

func parseMapdata(mapblock *MapBlock, data []byte) (int, error) {
	r := bytes.NewReader(data)

	cr := new(CountedReader)
	cr.Reader = r

	z, err := zlib.NewReader(cr)
	if err != nil {
		return 0, err
	}

	defer z.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, z)

	if buf.Len() != 16384 {
		return 0, errors.New("Mapdata length invalid: " + strconv.Itoa(buf.Len()))
	}

	rawdata := buf.Bytes()

	mapd := MapData{
		ContentId: make([]int, 4096),
		Param1:    make([]int, 4096),
		Param2:    make([]int, 4096),
	}
	mapblock.Mapdata = &mapd

	for i := 0; i < 4096; i++ {
		mapd.ContentId[i] = readU16(rawdata, i*2)
		mapd.Param1[i] = readU8(rawdata, (i*2)+2)
		mapd.Param2[i] = readU8(rawdata, (i*2)+3) //TODO: last item has wrong value
	}

	return cr.Count, nil
}
