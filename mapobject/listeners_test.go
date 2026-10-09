package mapobject

import (
	"testing"

	"github.com/minetest-go/mapparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPoiBlock(t *testing.T) {
	mb := newTestBlock("mapserver:poi", 1, 2, 3)
	setMeta(mb, 1, 2, 3, map[string]string{"name": "home", "owner": "singleplayer", "url": "http://x"})

	o := (&PoiBlock{Color: "red"}).onMapObject(testMBPos, 1, 2, 3, mb)
	require.NotNil(t, o)
	assert.Equal(t, "poi", o.Type)
	assert.Equal(t, 16+1, o.X)
	assert.Equal(t, -32+2, o.Y)
	assert.Equal(t, 48+3, o.Z)
	assert.Equal(t, "home", o.Attributes["name"])
	assert.Equal(t, "singleplayer", o.Attributes["owner"])
	assert.Equal(t, "red", o.Attributes["color"])
}

func TestProtectorBlock(t *testing.T) {
	mb := newTestBlock("protector:protect", 0, 0, 0)
	setMeta(mb, 0, 0, 0, map[string]string{"owner": "a", "members": "b c"})

	o := (&ProtectorBlock{}).onMapObject(testMBPos, 0, 0, 0, mb)
	require.NotNil(t, o)
	assert.Equal(t, "protector", o.Type)
	assert.Equal(t, "a", o.Attributes["owner"])
	assert.Equal(t, "b c", o.Attributes["members"])
}

func TestLabelBlock(t *testing.T) {
	mb := newTestBlock("mapserver:label", 0, 0, 0)
	setMeta(mb, 0, 0, 0, map[string]string{"text": "hi", "size": "5", "direction": "n", "owner": "o", "color": "#fff"})

	o := (&LabelBlock{}).onMapObject(testMBPos, 0, 0, 0, mb)
	require.NotNil(t, o)
	assert.Equal(t, "label", o.Type)
	assert.Equal(t, "hi", o.Attributes["text"])
	assert.Equal(t, "5", o.Attributes["size"])
	assert.Equal(t, "#fff", o.Attributes["color"])
}

func TestSignBlock(t *testing.T) {
	mb := newTestBlock("default:sign_wall_wood", 4, 4, 4)
	setMeta(mb, 4, 4, 4, map[string]string{"text": "hello"})

	o := (&SignBlock{Material: "wood"}).onMapObject(testMBPos, 4, 4, 4, mb)
	require.NotNil(t, o)
	assert.Equal(t, "sign", o.Type)
	assert.Equal(t, "hello", o.Attributes["display_text"])
	assert.Equal(t, "wood", o.Attributes["material"])
}

func TestDecodeUText(t *testing.T) {
	assert.Equal(t, "sign1\nA\na", decodeUText("return {115,105,103,110,49,10,65,10,97}"))
	assert.Equal(t, "", decodeUText("return {}"))
	assert.Equal(t, "", decodeUText(""))
	// out-of-range values are skipped
	assert.Equal(t, "AB", decodeUText("{65,300,66}"))
}

func TestMclSignBlock(t *testing.T) {
	mb := newTestBlock("mcl_signs:wall_sign", 0, 0, 0)
	setMeta(mb, 0, 0, 0, map[string]string{"utext": "return {104,105}"})

	o := (&MclSignBlock{Material: "oak"}).onMapObject(testMBPos, 0, 0, 0, mb)
	require.NotNil(t, o)
	assert.Equal(t, "hi", o.Attributes["display_text"])
	assert.Equal(t, "oak", o.Attributes["material"])
}

func TestTravelnetBlock(t *testing.T) {
	mb := newTestBlock("travelnet:travelnet", 2, 2, 2)
	setMeta(mb, 2, 2, 2, map[string]string{"owner": "o", "station_name": "Home", "station_network": "net"})

	o := (&TravelnetBlock{}).onMapObject(testMBPos, 2, 2, 2, mb)
	require.NotNil(t, o)
	assert.Equal(t, "travelnet", o.Type)
	assert.Equal(t, "Home", o.Attributes["station_name"])
	assert.Equal(t, "net", o.Attributes["station_network"])
	assert.Equal(t, "travelnet:travelnet", o.Attributes["nodename"])
}

func TestTravelnetBlockPrivateIgnored(t *testing.T) {
	mb := newTestBlock("travelnet:travelnet", 2, 2, 2)
	setMeta(mb, 2, 2, 2, map[string]string{"station_name": "(P)secret"})

	assert.Nil(t, (&TravelnetBlock{}).onMapObject(testMBPos, 2, 2, 2, mb))
}

func TestBonesBlock(t *testing.T) {
	tests := []struct {
		name  string
		meta  map[string]string
		owner string
	}{
		{"owner", map[string]string{"owner": "a"}, "a"},
		{"_owner fallback", map[string]string{"_owner": "b"}, "b"},
		{"unknown", map[string]string{}, "unknown"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mb := newTestBlock("bones:bones", 0, 0, 0)
			setMeta(mb, 0, 0, 0, tc.meta)
			setInv(mb, 0, 0, 0, "main",
				&mapparser.Item{Name: "default:stone", Count: 3},
				&mapparser.Item{Name: "default:dirt", Count: 4},
				&mapparser.Item{})

			o := (&BonesBlock{}).onMapObject(testMBPos, 0, 0, 0, mb)
			require.NotNil(t, o)
			assert.Equal(t, "bones", o.Type)
			assert.Equal(t, tc.owner, o.Attributes["owner"])
			assert.Equal(t, "7", o.Attributes["item_count"])
		})
	}
}

func TestBonesBlockNoInventory(t *testing.T) {
	mb := newTestBlock("bones:bones", 0, 0, 0)
	assert.Nil(t, (&BonesBlock{}).onMapObject(testMBPos, 0, 0, 0, mb))
}

func TestSmartShopBlock(t *testing.T) {
	newShop := func(shoptype string) *mapparser.MapBlock {
		mb := newTestBlock("smartshop:shop", 0, 0, 0)
		setMeta(mb, 0, 0, 0, map[string]string{"owner": "seller", "type": shoptype})
		setInv(mb, 0, 0, 0, "main",
			&mapparser.Item{Name: "default:apple", Count: 10},
			&mapparser.Item{Name: "default:apple", Count: 11},
			&mapparser.Item{Name: "default:stone", Count: 99})
		setInv(mb, 0, 0, 0, "pay1", &mapparser.Item{Name: "default:gold_ingot", Count: 2})
		setInv(mb, 0, 0, 0, "give1", &mapparser.Item{Name: "default:apple", Count: 5})
		// slot 2 has no give-item and is skipped
		setInv(mb, 0, 0, 0, "pay2", &mapparser.Item{Name: "default:gold_ingot", Count: 1})
		setInv(mb, 0, 0, 0, "give2")
		return mb
	}

	t.Run("normal shop", func(t *testing.T) {
		list := (&SmartShopBlock{}).onMapObject(testMBPos, 0, 0, 0, newShop("1"))
		require.Len(t, list, 1)
		o := list[0]
		assert.Equal(t, "shop", o.Type)
		assert.Equal(t, "smartshop", o.Attributes["type"])
		assert.Equal(t, "seller", o.Attributes["owner"])
		assert.Equal(t, "default:gold_ingot", o.Attributes["in_item"])
		assert.Equal(t, "2", o.Attributes["in_count"])
		assert.Equal(t, "default:apple", o.Attributes["out_item"])
		assert.Equal(t, "5", o.Attributes["out_count"])
		// 21 apples / 5 per trade
		assert.Equal(t, "4", o.Attributes["stock"])
	})

	t.Run("creative shop has infinite stock", func(t *testing.T) {
		list := (&SmartShopBlock{}).onMapObject(testMBPos, 0, 0, 0, newShop("0"))
		require.Len(t, list, 1)
		assert.Equal(t, "199", list[0].Attributes["stock"])
	})

	t.Run("empty main inventory", func(t *testing.T) {
		mb := newShop("1")
		setInv(mb, 0, 0, 0, "main")
		assert.Empty(t, (&SmartShopBlock{}).onMapObject(testMBPos, 0, 0, 0, mb))
	})
}

func TestFancyVend(t *testing.T) {
	newVend := func(nodename string) *mapparser.MapBlock {
		mb := newTestBlock(nodename, 0, 0, 0)
		setMeta(mb, 0, 0, 0, map[string]string{
			"owner":    "seller",
			"settings": "return {input_item_qty=3, output_item_qty=2}",
		})
		setInv(mb, 0, 0, 0, "wanted_item", &mapparser.Item{Name: "default:gold_ingot", Count: 1})
		setInv(mb, 0, 0, 0, "given_item", &mapparser.Item{Name: "default:apple", Count: 1})
		setInv(mb, 0, 0, 0, "main",
			&mapparser.Item{Name: "default:apple", Count: 4},
			&mapparser.Item{Name: "default:stone", Count: 4})
		return mb
	}

	t.Run("normal vendor", func(t *testing.T) {
		o := (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, newVend("fancy_vend:player_vendor"))
		require.NotNil(t, o)
		assert.Equal(t, "shop", o.Type)
		assert.Equal(t, "fancyvend", o.Attributes["type"])
		assert.Equal(t, "default:gold_ingot", o.Attributes["in_item"])
		assert.Equal(t, "3", o.Attributes["in_count"])
		assert.Equal(t, "default:apple", o.Attributes["out_item"])
		assert.Equal(t, "2", o.Attributes["out_count"])
		assert.Equal(t, "2", o.Attributes["stock"])
	})

	t.Run("admin vendor has fixed stock", func(t *testing.T) {
		o := (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, newVend("fancy_vend:admin_vendor"))
		require.NotNil(t, o)
		assert.Equal(t, "499", o.Attributes["stock"])
	})

	t.Run("missing inventories", func(t *testing.T) {
		mb := newVend("fancy_vend:player_vendor")
		delete(mb.Metadata.GetInventoryMapAtPos(0, 0, 0), "wanted_item")
		assert.Nil(t, (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, mb))
	})

	t.Run("empty wanted item", func(t *testing.T) {
		mb := newVend("fancy_vend:player_vendor")
		setInv(mb, 0, 0, 0, "wanted_item", &mapparser.Item{})
		assert.Nil(t, (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, mb))
	})

	t.Run("invalid settings", func(t *testing.T) {
		mb := newVend("fancy_vend:player_vendor")
		setMeta(mb, 0, 0, 0, map[string]string{"settings": "not lua {{{"})
		assert.Nil(t, (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, mb))
	})

	t.Run("settings without quantities", func(t *testing.T) {
		mb := newVend("fancy_vend:player_vendor")
		setMeta(mb, 0, 0, 0, map[string]string{"settings": "return {}"})
		assert.Nil(t, (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, mb))
	})

	t.Run("quantities below 1 are clamped", func(t *testing.T) {
		mb := newVend("fancy_vend:player_vendor")
		setMeta(mb, 0, 0, 0, map[string]string{"settings": "return {input_item_qty=0, output_item_qty=0}"})
		o := (&FancyVend{}).onMapObject(testMBPos, 0, 0, 0, mb)
		require.NotNil(t, o)
		assert.Equal(t, "1", o.Attributes["in_count"])
		assert.Equal(t, "1", o.Attributes["out_count"])
	})
}

func TestSmartShopBlockNoMainInventory(t *testing.T) {
	mb := newTestBlock("smartshop:shop", 0, 0, 0)
	assert.Empty(t, (&SmartShopBlock{}).onMapObject(testMBPos, 0, 0, 0, mb))
}
