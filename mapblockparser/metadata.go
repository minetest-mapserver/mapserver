package mapblockparser

import (
	"compress/zlib"
	"bytes"
	"errors"
	"strconv"
	"io"
	log "github.com/sirupsen/logrus"
)


func readU16(data []byte, offset int) int {
	return (int(data[offset]) << 8) | int(data[offset + 1])
}

func readU32(data []byte, offset int){
}

func parseMetadata(mapblock *MapBlock, data []byte) (int, error) {
	log.WithFields(log.Fields{
		"data-length": len(data),
	}).Debug("Parsing metadata")

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

	metadata := buf.Bytes()

	offset := 0
	version := metadata[offset]

	if version != 2 {
		return 0, errors.New("Wrong metadata version: " + strconv.Itoa(int(version)))
	}

	offset++;
	count := readU16(metadata, offset)

	log.WithFields(log.Fields{
		"count": count,
		"version": version,
	}).Debug("Parsed metadata-header")

	offset+=2

	log.Println("Metadata", buf.String())//XXX

	for i := 0; i < count; i++ {
		position := readU16(metadata, offset);
		log.Println("MD item", i, position)//XXX

		offset+=2

		//TODO
	}

	return cr.Count, nil
}

