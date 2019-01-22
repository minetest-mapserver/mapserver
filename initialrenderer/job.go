package initialrenderer

import (
	"mapserver/app"
	"mapserver/coords"
	"time"

	"github.com/sirupsen/logrus"
)

func Job(ctx *app.App) {

	fields := logrus.Fields{}
	logrus.WithFields(fields).Info("Starting initial rendering")

	rstate := ctx.Config.RenderState

	lastcoords := coords.NewMapBlockCoords(rstate.LastX, rstate.LastY, rstate.LastZ)

	for true {
		start := time.Now()

		hasMore, newlastcoords, mblist, err := ctx.BlockAccessor.FindLegacyMapBlocks(lastcoords, ctx.Config.InitialRenderingFetchLimit, ctx.Config.Layers)

		if err != nil {
			panic(err)
		}

		if len(mblist) == 0 && !hasMore {
			logrus.Info("Initial rendering complete")
			rstate.InitialRun = false
			ctx.Config.Save()

			break
		}

		lastcoords = *newlastcoords

		//Render zoom 12
		for _, mb := range mblist {
			//zoom 13
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			//zoom 12
			tc = tc.GetZoomedOutTile()

			fields = logrus.Fields{
				"X":       tc.X,
				"Y":       tc.Y,
				"Zoom":    tc.Zoom,
				"LayerId": tc.LayerId,
			}
			logrus.WithFields(fields).Debug("Dispatching tile rendering (z12)")

			ctx.Objectdb.RemoveTile(tc)
			_, err = ctx.Tilerenderer.Render(tc, 2)
			if err != nil {
				panic(err)
			}
		}

		//Render zoom 11-1
		for _, mb := range mblist {
			//13
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			//12
			tc = tc.GetZoomedOutTile()

			for tc.Zoom > 1 {
				//11-1
				tc = tc.GetZoomedOutTile()

				fields = logrus.Fields{
					"X":       tc.X,
					"Y":       tc.Y,
					"Zoom":    tc.Zoom,
					"LayerId": tc.LayerId,
				}
				logrus.WithFields(fields).Debug("Dispatching tile rendering (z11-1)")

				ctx.Objectdb.RemoveTile(tc)
				_, err = ctx.Tilerenderer.Render(tc, 1)
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

		t := time.Now()
		elapsed := t.Sub(start)

		fields = logrus.Fields{
			"count":   len(mblist),
			"X":       lastcoords.X,
			"Y":       lastcoords.Y,
			"Z":       lastcoords.Z,
			"elapsed": elapsed,
		}
		logrus.WithFields(fields).Info("Initial rendering")
	}
}
