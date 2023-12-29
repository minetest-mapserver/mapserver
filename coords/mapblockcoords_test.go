package coords

import (
	"mapserver/types"
	"testing"
)

func TestNewMapBlockCoordsFromBlock(t *testing.T) {
	c := types.NewMapBlockCoordsFromBlock(1, 1, 1)

	if c.X != 0 || c.Y != 0 || c.Z != 0 {
		t.Fatal("mismatch", c)
	}

	c = types.NewMapBlockCoordsFromBlock(16, 1, 1)

	if c.X != 1 || c.Y != 0 || c.Z != 0 {
		t.Fatal("mismatch", c)
	}

	c = types.NewMapBlockCoordsFromBlock(16, 1, -1)

	if c.X != 1 || c.Y != 0 || c.Z != -1 {
		t.Fatal("mismatch", c)
	}

}
