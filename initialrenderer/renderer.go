package initialrenderer

import (
  "mapserver/tilerenderer"
  "mapserver/layerconfig"
  "mapserver/coords"
  "github.com/sirupsen/logrus"
  "time"

)

func worker(tr *tilerenderer.TileRenderer, jobs <-chan coords.TileCoords){
  for coord := range(jobs){
    tr.Render(coord)
  }
}

func Render(tr *tilerenderer.TileRenderer,
  layers []layerconfig.Layer){

    start := time.Now()
    complete_count := 512*512
    current_count := 0
    perf_count := 0

    jobs := make(chan coords.TileCoords, 100)

    go worker(tr, jobs)
    go worker(tr, jobs)
    go worker(tr, jobs)
    go worker(tr, jobs)

    for _, layer := range(layers) {

    	//zoom 10 iterator
    	for x := -255; x<256; x++ {
    		for y := -255; y<256; y++ {
    			tc := coords.NewTileCoords(x,y,10,layer.Id)
          jobs <- tc
          current_count++
          perf_count++

          if time.Now().Sub(start).Seconds() > 2 {
            start = time.Now()
            progress := float64(current_count) / float64(complete_count) * 100

            fields := logrus.Fields{
              "x": x,
              "y": y,
              "progress%": progress,
              "layer": layer.Name,
              "perf": perf_count,
            }

            perf_count = 0
            logrus.WithFields(fields).Info("Initial render progress")
          }
    		}
    	}


    }

    close(jobs)

}


// zoom:1 == length=1
// zoom:2 == length=2
// zoom:3 == length=4
// zoom:4 == length=8
// zoom:5 == length=16
// zoom:6 == length=32
// zoom:7 == length=64
// zoom:8 == length=128
// zoom:9 == length=256
// zoom:10 == length=512
// zoom:11 == length=1024
// zoom:12 == length=2048
// zoom:13 == length=4096
