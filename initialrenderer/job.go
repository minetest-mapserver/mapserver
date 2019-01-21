package initialrenderer

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
)

func Job(ctx *app.App) {

	fields := logrus.Fields{}
	logrus.WithFields(fields).Info("Starting initial rendering")

	rstate := ctx.Config.RenderState

	lastcoords := coords.NewMapBlockCoords(rstate.LastX, rstate.LastY, rstate.LastZ)

	for true {
		newlastcoords, mblist, err := ctx.BlockAccessor.FindLegacyMapBlocks(lastcoords, 1000)

		if err != nil {
			panic(err)
		}

		if len(mblist) == 0 {
			logrus.Info("Initial rendering complete")
			rstate.InitialRun = false
			ctx.Config.Save()

			break
		}

		lastcoords = *newlastcoords

		for _, mb := range mblist {
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos)

			_, err = ctx.Tilerenderer.Render(tc)
			if err != nil {
				panic(err)
			}
		}

		//Save current positions of initial run
		rstate.LastX = lastcoords.X
		rstate.LastY = lastcoords.Y
		rstate.LastZ = lastcoords.Z
		ctx.Config.Save()

		fields = logrus.Fields{
			"count": len(mblist),
			"X":     lastcoords.X,
			"Y":     lastcoords.Y,
			"Z":     lastcoords.Z,
		}
		logrus.WithFields(fields).Info("Initial rendering")
	}
}
