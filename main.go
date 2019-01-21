package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mapserver/app"
	"mapserver/initialrenderer"
	"mapserver/params"
	"mapserver/tileupdate"
	"mapserver/web"
)

func main() {
	//Parse command line

	p := params.Parse()

	if p.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

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
	cfg, err := app.ParseConfig(app.ConfigFile)
	if err != nil {
		panic(err)
	}

	//write back config with all values
	err = cfg.Save()
	if err != nil {
		panic(err)
	}

	//setup app context
	ctx, err := app.Setup(p, cfg)

	if err != nil {
		//error case
		panic(err)
	}

	//run initial rendering
	if ctx.Config.EnableInitialRendering && ctx.Config.RenderState.InitialRun {
		go initialrenderer.Job(ctx)
	}

	//Incremental update
	if ctx.Config.EnableIncrementalUpdate {
		go tileupdate.Job(ctx)
	}

	//Start http server
	web.Serve(ctx)

}
