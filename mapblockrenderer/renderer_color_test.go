package mapblockrenderer

import (
	"image/color"
	"mapserver/mapblockaccessor"
	"mapserver/types"
	"sync"
	"testing"
	"time"

	"github.com/luanti-golang/colormapping"
	"github.com/luanti-golang/mapparser"
)

const (
	testPaletteId = 2
	testUnknownId = 3
	paletteNode   = "unifiedbricks:brickblock"
)

func TestRenderColors(t *testing.T) {
	c := colormapping.NewColorMapping()
	_, err := c.LoadBytes([]byte("default:stone 100 100 100 128\n"))
	if err != nil {
		t.Fatal(err)
	}

	mb := newTestMapblock()
	mb.BlockMapping[testPaletteId] = paletteNode
	mb.BlockMapping[testUnknownId] = "unknown:node"
	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			setTestNode(mb, x, 0, z, testStoneId)
		}
	}
	// param2 colored nodes
	setTestNode(mb, 3, 1, 3, testPaletteId)
	mb.Param2[mapparser.GetNodePos(3, 1, 3)] = 10
	setTestNode(mb, 7, 1, 7, testPaletteId)
	mb.Param2[mapparser.GetNodePos(7, 1, 7)] = 200
	// node without color: the stone below is visible
	setTestNode(mb, 10, 1, 10, testUnknownId)

	accessor := mapblockaccessor.NewMapBlockAccessor(nil, time.Minute, 100)
	pos := types.NewMapBlockCoords(0, 0, 0)
	accessor.Update(pos, mb)
	accessor.Update(types.NewMapBlockCoords(-1, 0, 0), nil)
	accessor.Update(types.NewMapBlockCoords(0, 0, 1), nil)

	r := NewMapBlockRenderer(accessor, c)
	r.enableShadow = false

	// concurrent renders share the color cache (run with -race)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := r.Render(pos, pos); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()

	img, err := r.Render(pos, pos)
	if err != nil {
		t.Fatal(err)
	}

	nodeColor := func(x, z int) color.NRGBA {
		return img.NRGBAAt(x*IMG_SCALE, (15-z)*IMG_SCALE)
	}
	paletteColor := func(param2 int) color.NRGBA {
		pc := c.GetColor(paletteNode, param2)
		return color.NRGBA{R: pc.R, G: pc.G, B: pc.B, A: 255}
	}
	stone := color.NRGBA{R: 100, G: 100, B: 100, A: 255}

	if got := nodeColor(3, 3); got != paletteColor(10) {
		t.Errorf("palette node param2=10: got %v, want %v", got, paletteColor(10))
	}
	if got := nodeColor(7, 7); got != paletteColor(200) {
		t.Errorf("palette node param2=200: got %v, want %v", got, paletteColor(200))
	}
	if paletteColor(10) == paletteColor(200) {
		t.Error("test palette colors should differ")
	}
	if got := nodeColor(10, 10); got != stone {
		t.Errorf("unknown node: got %v, want stone %v", got, stone)
	}

	// the shared colormapping color is not modified (alpha clamped on a copy)
	if a := c.GetColor("default:stone", 0).A; a != 128 {
		t.Errorf("colormapping color modified: alpha %d", a)
	}
}
