package util

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/minetest-go/mapparser"
)

func SerializeMapdata(mapdata *mapparser.MapData) (string, error) {
	mapbytes := make([]byte, 4096*4) // 16^3: nodeid's (16 bit), param1 (8 bit), param2 (8 bit)

	for i := 0; i < 4096; i++ {
		binary.BigEndian.PutUint16(mapbytes[i*2:], uint16(mapdata.ContentId[i]))
		mapbytes[(4096*2)+i] = byte(mapdata.Param1[i])
		mapbytes[(4096*3)+i] = byte(mapdata.Param2[i])
	}

	buf := bytes.NewBuffer([]byte{})
	zw := zlib.NewWriter(buf)
	_, err := zw.Write(mapbytes)
	if err != nil {
		return "", fmt.Errorf("zlib write error: %v", err)
	}
	err = zw.Close()
	if err != nil {
		return "", fmt.Errorf("zlib close error: %v", err)
	}

	return base64.RawStdEncoding.EncodeToString(buf.Bytes()), nil
}
