package mapblockparser

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"strconv"
)

//TODO: mapdata struct with accessors

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

	mapblock.Mapdata = buf.Bytes()

	return cr.Count, nil
}
