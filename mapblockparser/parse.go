package mapblockparser

import (
	"errors"
	"compress/zlib"
	"bytes"
	"io"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type MapBlock struct {
	Version byte
	Underground bool
	Mapdata []byte
}

func readU16(data []byte, offset int){
}

func readU32(data []byte, offset int){
}


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

func parseMetadata(mapblock *MapBlock, data []byte) (int, error) {
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

	log.Println("Metadata length ", buf.Len(), buf.String())

	return cr.Count, nil
}

func Parse(data []byte) (*MapBlock, error) {
	mapblock := MapBlock{}
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	log.Println("data-length: ", len(data))

	offset := 0

	// version
	mapblock.Version = data[0]

	//flags
	flags := data[1]
	mapblock.Underground = (flags & 0x01) == 0x01

	//mapdata (blocks)
	offset = 6

	//metadata
	count, err := parseMapdata(&mapblock, data[offset:])
	if err != nil {
		return nil, err
	}

	log.Println("Mapdata length: ", count)

	offset += count

	log.Println("New offset: ", offset)

	count, err = parseMetadata(&mapblock, data[offset:])
	if err != nil {
		return nil, err
	}

	return &mapblock, nil
}

