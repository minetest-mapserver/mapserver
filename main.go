package main

import (
	"flag"
	"fmt"
	"image/png"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/mapblockaccessor"
	"mapserver/mapblockrenderer"
	"mapserver/params"
	"mapserver/worldconfig"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	Version = "2.0-DEV"
)

type JobData struct {
	pos1, pos2 coords.MapBlockCoords
	x, z       int
}

func worker(r *mapblockrenderer.MapBlockRenderer, jobs <-chan JobData) {
	for d := range jobs {
		img, _ := r.Render(d.pos1, d.pos2)

		if img != nil {
			f, _ := os.Create(fmt.Sprintf("output/image_%d_%d.png", d.x, d.z))
			start := time.Now()
			png.Encode(f, img)
			f.Close()
			t := time.Now()
			elapsed := t.Sub(start)
			logrus.WithFields(logrus.Fields{"elapsed": elapsed}).Debug("Encoding completed")
		}
	}
}

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

	jobs := make(chan JobData, 100)
	go worker(&r, jobs)
	go worker(&r, jobs)
	go worker(&r, jobs)
	go worker(&r, jobs)

	from := -500
	to := 500

	logrus.WithFields(logrus.Fields{"from": from, "to": to}).Info("Starting rendering")

	for x := from; x < to; x++ {
		for z := from; z < to; z++ {
			pos1 := coords.NewMapBlockCoords(x, 10, z)
			pos2 := coords.NewMapBlockCoords(x, -1, z)

			jobs <- JobData{pos1: pos1, pos2: pos2, x: x, z: z}
		}
	}

	close(jobs)
}
