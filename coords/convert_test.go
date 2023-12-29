package coords

import (
	"mapserver/types"
	"testing"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"prefix": "coords/convert_test"})
}

func TestMinCoord(t *testing.T) {
	c := types.NewMapBlockCoords(types.MinCoord, types.MinCoord, types.MinCoord)
	pc := CoordToPlain(c)

	log.WithFields(logrus.Fields{"coords": c, "plain": pc, "plain-1": pc - 1}).Info("TestMinCoord")
	if pc != MinPlainCoord {
		t.Fatal("no min match")
	}
}

func testCoordConvert(t *testing.T, mb *types.MapBlockCoords) {
	log.WithFields(logrus.Fields{"coords": mb}).Info("MapblockCoords")

	p := CoordToPlain(mb)
	log.WithFields(logrus.Fields{"plain": p}).Info("MapblockCoords")

	mb2 := PlainToCoord(p)
	log.WithFields(logrus.Fields{"coords2": mb2}).Info("MapblockCoords")

	if mb.X != mb2.X {
		t.Fatal("X mismatch")
	}

	if mb.Y != mb2.Y {
		t.Fatal("Y mismatch")
	}

	if mb.Z != mb2.Z {
		t.Fatal("Z mismatch")
	}

}

func TestZeroCoord(t *testing.T) {
	testCoordConvert(t, types.NewMapBlockCoords(0, 0, 0))
}
