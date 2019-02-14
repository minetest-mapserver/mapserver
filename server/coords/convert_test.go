package coords

import (
	"testing"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"prefix": "coords/convert_test"})
}

func TestMinCoord(t *testing.T){
	c := NewMapBlockCoords(MinCoord, MinCoord, MinCoord)
	pc := CoordToPlain(c)

	log.WithFields(logrus.Fields{"coords": c, "plain": pc}).Info("TestMinCoord")
	if pc != MinPlainCoord {
		t.Fatal("no min match")
	}
}

func testCoordConvert(t *testing.T, mb *MapBlockCoords) {
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

func TestZeorCoord(t *testing.T) {
	testCoordConvert(t, NewMapBlockCoords(0, 0, 0))
}
