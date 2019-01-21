package initialrenderer

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"time"
)

func Job(ctx *app.App) {

	fields := logrus.Fields{}
	logrus.WithFields(fields).Info("Starting initial rendering")

	lastcoords := coords.NewMapBlockCoords(coords.MinCoord, coords.MinCoord, coords.MinCoord)

	for true {
		newlastcoords, mblist, err := ctx.BlockAccessor.FindLegacyMapBlocks(lastcoords, 1000)

		if err != nil {
			panic(err)
		}

		lastcoords = *newlastcoords

		if len(mblist) == 1 {
			logrus.Info("Initial rendering complete")
			break
		}

		//for _, mb := range mblist {
		//}

		fields = logrus.Fields{
			"count": len(mblist),
			"X":  lastcoords.X,
			"Y":  lastcoords.Y,
			"Z":  lastcoords.Z,
		}
		logrus.WithFields(fields).Info("Initial rendering")

		time.Sleep(100 * time.Millisecond)
	}
}
