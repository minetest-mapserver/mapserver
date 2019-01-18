package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/initialrenderer"
	"mapserver/layerconfig"
	"mapserver/params"
	"mapserver/web"

	"fmt"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	//Parse command line

	p := params.Parse()

	if p.Help {
		params.PrintHelp()
		return
	}

	if p.Version {
		fmt.Print("Mapserver version: ")
		fmt.Println(app.Version)
		return
	}

	//parse Config
	cfg, err := app.ParseConfig("mapserver.json")
	if err != nil {
		panic(err)
	}

	if p.Dumpconfig {
		str, err := json.MarshalIndent(cfg, "", "	")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(str))
		return
	}

	//setup app context
	ctx, err := app.Setup(p, cfg)

	if err != nil {
		//error case
		panic(err)
	}

	//run initial rendering
	if ctx.Config.EnableInitialRendering {
		go initialrenderer.Render(ctx.Tilerenderer, layerconfig.DefaultLayers)
	}

	//Start http server
	web.Serve(ctx)

}
