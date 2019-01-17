package main

import (
	"flag"
	"fmt"
	"time"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/params"
	"mapserver/worldconfig"
	"os"
	"github.com/sirupsen/logrus"
)

const (
	Version = "2.0-DEV"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	p := params.Parse()

	if p.Help {
		flag.PrintDefaults()
		return
	}

	if p.Version {
		fmt.Print("Mapserver version: ")
		fmt.Println(Version)
		return
	}

	worldcfg := worldconfig.Parse(p.Worlddir + "world.mt")
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mapserver")

	if worldcfg.Backend != worldconfig.BACKEND_SQLITE3 {
		panic("no sqlite3 backend found!")
	}

	a, err := db.NewSqliteAccessor("map.sqlite")
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	cache := mapblockaccessor.NewMapBlockAccessor(a)
	c := colormapping.NewColorMapping()
	err = c.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		panic(err)
	}

	r := mapblockrenderer.NewMapBlockRenderer(cache, c)
	os.Mkdir("output", 0755)

	results := make(chan mapblockrenderer.JobResult, 100)
	jobs := make(chan mapblockrenderer.JobData, 100)

	for i := 0; i<3; i++ {
		go mapblockrenderer.Worker(&r, jobs, results)
	}

	os.Mkdir("output", 0755)

	go func() {
		for result := range results {
			if result.Data.Len() == 0 {
				continue
			}

			tc := coords.GetTileCoordsFromMapBlock(result.Job.Pos1)
			f, _ := os.Create(fmt.Sprintf("output/image_%d_%d.png", tc.X, tc.Y))
			result.Data.WriteTo(f)
			f.Close()
		}
	}()

	from := coords.MinCoord
	to := coords.MaxCoord

	logrus.WithFields(logrus.Fields{"from": from, "to": to}).Info("Starting rendering")

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
				logrus.WithFields(logrus.Fields{"x": x, "z": z, "progress%": progress}).Info("Render progress")
			}
		}
	}

	close(jobs)
	defer close(results)
}
