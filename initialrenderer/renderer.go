package initialrenderer

import (
  "mapserver/mapblockrenderer"
  "mapserver/tiledb"
  "mapserver/coords"
  "github.com/sirupsen/logrus"
  "time"

)

func Render(renderer *mapblockrenderer.MapBlockRenderer,
  tdb tiledb.DBAccessor){

    results := make(chan mapblockrenderer.JobResult, 100)
  	jobs := make(chan mapblockrenderer.JobData, 100)

  	for i := 0; i<3; i++ {
  		go mapblockrenderer.Worker(renderer, jobs, results)
  	}

    go func() {
      for result := range results {
        tc := coords.GetTileCoordsFromMapBlock(result.Job.Pos1)
        tile := tiledb.Tile{Pos: tc, Data: result.Data.Bytes(), Mtime: time.Now().Unix()}
        tdb.SetTile(&tile)
      }
    }()

    from := coords.MinCoord
    to := coords.MaxCoord

    start := time.Now()
    complete_count := (to - from) * (to - from)
    current_count := 0

    for x := from; x < to; x++ {
      for z := from; z < to; z++ {
        pos1 := coords.NewMapBlockCoords(x, 10, z)
        pos2 := coords.NewMapBlockCoords(x, -1, z)

        jobs <- mapblockrenderer.JobData{Pos1: pos1, Pos2: pos2}
        current_count++

        if time.Now().Sub(start).Seconds() > 2 {
          start = time.Now()
          progress := float64(current_count) / float64(complete_count) * 100
          logrus.WithFields(logrus.Fields{"x": x, "z": z, "progress%": progress}).Info("Initial render progress")
        }
      }
    }

    close(jobs)
    defer close(results)

}
