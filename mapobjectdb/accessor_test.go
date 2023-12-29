package mapobjectdb

import (
	"fmt"
	"mapserver/types"
	"testing"
)

func TestNewMapBlockCoords(t *testing.T) {
	attrs := make(map[string]string)
	attrs["X"] = "y"

	pos := types.NewMapBlockCoords(1, 2, 3)
	fmt.Println(pos)

	obj := NewMapObject(pos, 10, 12, 14, "xy")
	fmt.Println(obj)

	if obj.X != 26 {
		t.Error("x coord off")
	}

	if obj.Y != 44 {
		t.Error("Y coord off")
	}

	if obj.Z != 62 {
		t.Error("Z coord off")
	}

}
