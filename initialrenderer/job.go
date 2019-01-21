package initialrenderer

import (
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapblockparser"
)

func Job(ctx *app.App) {

	fields := logrus.Fields{}
	logrus.WithFields(fields).Info("Starting initial rendering")

	rstate := ctx.Config.RenderState

	lastcoords := coords.NewMapBlockCoords(rstate.LastX, rstate.LastY, rstate.LastZ)

	for true {
		newlastcoords, mblist, err := ctx.BlockAccessor.FindLegacyMapBlocks(lastcoords, 10000)

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

		//only mapblocks with valid layer
		validmblist := make([]*mapblockparser.MapBlock, 0)

		//Invalidate zoom 12-1
		for _, mb := range mblist {
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			if tc == nil {
				continue
			}

			validmblist = append(validmblist, mb)

			for tc.Zoom > 1 {
				tc = tc.GetZoomedOutTile()
				ctx.Tiledb.RemoveTile(tc)
			}
		}

		//Render zoom 12-1
		for _, mb := range validmblist {
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)
			for tc.Zoom > 1 {
				tc = tc.GetZoomedOutTile()

				fields = logrus.Fields{
					"X":     tc.X,
					"Y":     tc.Y,
					"Zoom":     tc.Zoom,
					"LayerId": tc.LayerId,
				}
				logrus.WithFields(fields).Debug("Dispatching tile rendering")

				_, err = ctx.Tilerenderer.Render(tc)
				if err != nil {
					panic(err)
				}
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
			"validcount": len(validmblist),
		}
		logrus.WithFields(fields).Info("Initial rendering")
	}
}
