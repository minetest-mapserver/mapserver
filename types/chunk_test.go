package types

import (
	"testing"
)

func TestGetMapblockPosFromNode(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name  string
		args  args
		wantx int
		wanty int
		wantz int
	}{
		{"coords for node 0,0,0", args{0, 0, 0}, 0, 0, 0},
		{"coords for node 15,15,15", args{15, 15, 15}, 0, 0, 0},
		{"coords for node 16,16,16", args{16, 16, 16}, 1, 1, 1},
		{"coords for node 31,16,16", args{31, 16, 16}, 1, 1, 1},
		{"coords for node -15,-15,-15", args{-15, -15, -15}, -1, -1, -1},
		{"coords for node -1565,4,-386", args{-1565, 4, -386}, -98, 0, -25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, z := GetMapblockPosFromNode(tt.args.x, tt.args.y, tt.args.z)
			t.Logf("GetMapblockPosFromNode %v: %v, %v, %v", tt.args, x, y, z)
			if x != tt.wantx {
				t.Errorf("GetMapblockPosFromNode() x = %v, want %v", x, tt.wantx)
			}
			if y != tt.wanty {
				t.Errorf("GetMapblockPosFromNode() y = %v, want %v", y, tt.wanty)
			}
			if z != tt.wantz {
				t.Errorf("GetMapblockPosFromNode() z = %v, want %v", z, tt.wantz)
			}
		})
	}
}
