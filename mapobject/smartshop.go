package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"
	"math"
	"strconv"

	"github.com/minetest-go/mapparser"
)

type SmartShopBlock struct{}

func (this *SmartShopBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) []*mapobjectdb.MapObject {
	list := make([]*mapobjectdb.MapObject, 0)

	md := block.Metadata.GetMetadata(x, y, z)
	invMap := block.Metadata.GetInventoryMapAtPos(x, y, z)
	mainInv := invMap["main"]
	isCreative := md["type"] == "0"

	if mainInv.IsEmpty() {
		return list
	}

	for i := 1; i <= 4; i++ {
		payInvName := "pay" + strconv.Itoa(i)
		giveInvName := "give" + strconv.Itoa(i)

		pay := invMap[payInvName]
		give := invMap[giveInvName]

		if len(pay.Items) == 0 || len(give.Items) == 0 {
			continue
		}

		o := mapobjectdb.NewMapObject(mbpos, x, y, z, "shop")
		o.Attributes["type"] = "smartshop"
		o.Attributes["owner"] = md["owner"]

		in_item := pay.Items[0].Name
		in_count := math.Max(1, float64(pay.Items[0].Count))

		out_item := give.Items[0].Name
		out_count := math.Max(1, float64(give.Items[0].Count))

		stock := 0

		if isCreative {
			stock = 999

		} else {
			for _, item := range mainInv.Items {
				if item.Name == out_item {
					stock += item.Count
				}
			}

		}

		//multiples of out_count
		stock_factor := math.Floor(float64(stock) / float64(out_count))

		o.Attributes["in_item"] = in_item
		o.Attributes["in_count"] = strconv.Itoa(int(in_count))
		o.Attributes["out_item"] = out_item
		o.Attributes["out_count"] = strconv.Itoa(int(out_count))
		o.Attributes["stock"] = strconv.Itoa(int(stock_factor))

		list = append(list, o)
	}

	return list
}
