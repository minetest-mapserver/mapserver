package mapobject

import (
	"mapserver/luaparser"
	"mapserver/mapblockparser"
	"mapserver/mapobjectdb"
	"math"
	"strconv"

	"github.com/sirupsen/logrus"
)

type FancyVend struct{}

func (this *FancyVend) onMapObject(x, y, z int, block *mapblockparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)
	nodename := block.GetNodeName(x, y, z)
	invMap := block.Metadata.GetInventoryMapAtPos(x, y, z)
	parser := luaparser.New()

	isAdmin := false

	if nodename == "fancy_vend:admin_vendor" || nodename == "fancy_vend:admin_depo" {
		isAdmin = true
	}

	payInv := invMap["wanted_item"]
	giveInv := invMap["given_item"]
	mainInv := invMap["main"]

	if payInv == nil || giveInv == nil {
		return nil
	}

	if payInv.Items == nil || giveInv.Items == nil {
		return nil
	}

	if payInv.Items[0].IsEmpty() || giveInv.Items[0].IsEmpty() {
		return nil
	}

	settings, err := parser.ParseMap(md["settings"])
	if err != nil {
		fields := logrus.Fields{
			"x":   x,
			"y":   y,
			"z":   z,
			"pos": block.Pos,
			"err": err,
		}
		log.WithFields(fields).Error("Fancyvend setting error")
		return nil
	}

	if settings["input_item_qty"] == nil || settings["output_item_qty"] == nil {
		return nil
	}

	in_count := settings["input_item_qty"].(int)
	if in_count < 1 {
		in_count = 1
	}
	out_count := settings["output_item_qty"].(int)
	if out_count < 1 {
		out_count = 1
	}

	in_item := payInv.Items[0].Name
	out_item := giveInv.Items[0].Name

	stock := 0

	if isAdmin {
		stock = 999

	} else {
		for _, item := range mainInv.Items {
			if item.Name == out_item {
				stock += int(math.Max(1, float64(item.Count)))
			}
		}
	}

	stock_factor := int(float64(stock) / float64(out_count))

	o := mapobjectdb.NewMapObject(block.Pos, x, y, z, "shop")
	o.Attributes["owner"] = md["owner"]
	o.Attributes["type"] = "fancyvend"

	o.Attributes["in_item"] = in_item
	o.Attributes["in_count"] = strconv.Itoa(in_count)
	o.Attributes["out_item"] = out_item
	o.Attributes["out_count"] = strconv.Itoa(out_count)
	o.Attributes["stock"] = strconv.Itoa(stock_factor)

	return o
}
