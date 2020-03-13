package mapblockparser

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

/*
lua vm: https://github.com/yuin/gopher-lua
*/

const (
	INVENTORY_TERMINATOR = "EndInventory"
	INVENTORY_END        = "EndInventoryList"
	INVENTORY_START      = "List"
)

func readU16(data []byte, offset int) int {
	return (int(data[offset]) << 8) | int(data[offset+1])
}

func readU8(data []byte, offset int) int {
	return int(data[offset])
}

func readU32(data []byte, offset int) int {
	return int(data[offset])<<24 |
		int(data[offset+1])<<16 |
		int(data[offset+2])<<8 |
		int(data[offset+3])
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

	if cr.Count == 0 {
		return 0, errors.New("no data")
	}

	metadata := buf.Bytes()

	log.WithFields(logrus.Fields{"metadata-length": len(metadata)}).Trace("Parsing metadata")

	offset := 0
	version := metadata[offset]

	if version == 0 {
		//No data?
		return cr.Count, nil
	}

	offset++
	count := readU16(metadata, offset)

	offset += 2

	for i := 0; i < count; i++ {
		position := readU16(metadata, offset)
		pairsMap := mapblock.Metadata.GetPairsMap(position)

		offset += 2
		valuecount := readU32(metadata, offset)

		offset += 4
		for j := 0; j < valuecount; j++ {
			keyLength := readU16(metadata, offset)
			offset += 2

			key := string(metadata[offset : keyLength+offset])
			offset += keyLength

			valueLength := readU32(metadata, offset)
			offset += 4

			if len(metadata) <= valueLength+offset {
				return 0, errors.New("metadata too short: " + strconv.Itoa(len(metadata)) +
					", valuelength: " + strconv.Itoa(int(valueLength)))
			}

			value := string(metadata[offset : valueLength+offset])
			offset += valueLength

			pairsMap[key] = value

			priv := 0
			if version >= 2 { /* private tag doesn't exist in version=1 */
				priv = readU8(metadata, offset)
				offset++
			}

			/*
				if priv != 0 {
					// do something usefull
					logrus.Info("Private items in Inventory")
				}
			*/
		}

		var currentInventoryName *string
		var currentInventory *Inventory

		scanner := bufio.NewScanner(bytes.NewReader(metadata[offset:]))
		for scanner.Scan() {
			txt := scanner.Text()
			offset += len(txt) + 1

			log.WithFields(logrus.Fields{"txt": txt, "position": position}).Trace("Parsing inventory")

			if strings.HasPrefix(txt, INVENTORY_START) {
				pairs := strings.Split(txt, " ")
				currentInventoryName = &pairs[1]
				currentInventory = mapblock.Metadata.GetInventory(position, *currentInventoryName)
				currentInventory.Size = 0

			} else if txt == INVENTORY_END {
				currentInventoryName = nil
				currentInventory = nil
			} else if currentInventory != nil {
				//content
				if strings.HasPrefix(txt, "Item") {
					item := Item{}
					parts := strings.Split(txt, " ")

					if len(parts) >= 2 {
						item.Name = parts[1]
					}

					if len(parts) >= 3 {
						val, err := strconv.ParseInt(parts[2], 10, 32)
						if err != nil {
							return 0, err
						}
						item.Count = int(val)
					}

					if len(parts) >= 4 {
						val, err := strconv.ParseInt(parts[3], 10, 32)
						if err != nil {
							return 0, err
						}
						item.Count = int(val)
					}

					currentInventory.Items = append(currentInventory.Items, &item)
					currentInventory.Size += 1

				}

			} else if txt == INVENTORY_TERMINATOR {
				break

			} else {
				return 0, errors.New("Malformed inventory: " + txt)
			}
		}

		//TODO

	}

	return cr.Count, nil
}
