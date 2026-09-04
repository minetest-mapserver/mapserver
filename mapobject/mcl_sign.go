package mapobject

import (
    "regexp"
    "strconv"
    "strings"
	"mapserver/mapobjectdb"
	"mapserver/types"

	"github.com/minetest-go/mapparser"
)

var utextNumberRE = regexp.MustCompile(`\d+`)
func decodeUText(value string) string {
    // Extract all decimal numbers from:
    // return {115,105,103,110,49,10,65,10,97}
    matches := utextNumberRE.FindAllString(value, -1)
    if len(matches) == 0 {
        return ""
    }

    var b strings.Builder

    for _, match := range matches {
        n, err := strconv.Atoi(match)
        if err != nil || n < 0 || n > 255 {
            continue
        }

        b.WriteByte(byte(n))
    }

    return b.String()
}

type MclSignBlock struct {
	Material string
}

func (this *MclSignBlock) onMapObject(mbpos *types.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)
    displayText := decodeUText(md["utext"])

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "sign")
	o.Attributes["display_text"] = displayText
	o.Attributes["material"] = this.Material

	return o
}
