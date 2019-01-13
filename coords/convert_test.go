package coords

import (
	"github.com/sirupsen/logrus"
	"testing"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"prefix": "coords/convert_test"})
}

func testCoordConvert(t *testing.T, mb MapBlockCoords) {
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

func TestConvertPlainMapBlock(t *testing.T) {
	testCoordConvert(t, NewMapBlockCoords(10, 0, -10))
	testCoordConvert(t, NewMapBlockCoords(-2048, 2047, -10))
	testCoordConvert(t, NewMapBlockCoords(-3, 0, 2047)) //0...2047

}
