package mapblockparser

import (
	"errors"
	"strconv"
	log "github.com/sirupsen/logrus"
)



func Parse(data []byte) (*MapBlock, error) {
	mapblock := MapBlock{}
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	log.WithFields(log.Fields{
		"data-length": len(data),
	}).Debug("Parsing mapblock")

	offset := 0

	// version
	mapblock.Version = data[0]

	//flags
	flags := data[1]
	mapblock.Underground = (flags & 0x01) == 0x01

	content_width := data[4]
	params_width := data[4]

	if content_width != 2 {
		return nil, errors.New("content_width = " + strconv.Itoa(int(content_width)))
	}

	if params_width != 2 {
		return nil, errors.New("params_width = " + strconv.Itoa(int(params_width)))
	}

	//mapdata (blocks)
	offset = 6

	//metadata
	count, err := parseMapdata(&mapblock, data[offset:])
	if err != nil {
		return nil, err
	}

	offset += count

	count, err = parseMetadata(&mapblock, data[offset:])
	if err != nil {
		return nil, err
	}

	return &mapblock, nil
}

