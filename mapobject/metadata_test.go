package mapobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseItem(t *testing.T) {
	assert.Equal(t, &Item{Name: "default:stone", Count: 1}, parseItem("default:stone"))
	assert.Equal(t, &Item{Name: "default:stone", Count: 99}, parseItem("default:stone 99"))
	assert.Equal(t, &Item{Name: "default:pick_steel", Count: 1, Wear: 1234}, parseItem("default:pick_steel 1 1234"))
	assert.Equal(t, &Item{Name: "default:pick_steel", Count: 1, Wear: 1234}, parseItem("default:pick_steel 1 1234 \x01description\x02x\x03"))
}

func TestGetInventoryMap(t *testing.T) {
	mb := newTestBlock("default:chest", 1, 2, 3)
	mb.Inventory[1+2*16+3*256] = map[string][]string{
		"main":  {"", "default:stone 3", "", "default:dirt"},
		"empty": {"", ""},
	}

	invMap := getInventoryMap(mb, 1, 2, 3)
	require.NotNil(t, invMap["main"])
	// empty slots are skipped
	require.Len(t, invMap["main"].Items, 2)
	assert.Equal(t, "default:stone", invMap["main"].Items[0].Name)
	assert.Equal(t, 3, invMap["main"].Items[0].Count)
	assert.Equal(t, 1, invMap["main"].Items[1].Count)
	assert.False(t, invMap["main"].IsEmpty())

	require.NotNil(t, invMap["empty"])
	assert.True(t, invMap["empty"].IsEmpty())
	assert.Nil(t, invMap["missing"])

	// node without metadata
	assert.Empty(t, getInventoryMap(mb, 0, 0, 0))
	assert.Equal(t, "", getMetadata(mb, 0, 0, 0)["owner"])
}
