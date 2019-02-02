package main

import (
	"fmt"
	"mapserver/app"
	"mapserver/mapobject"
	"mapserver/params"
	"mapserver/tilerendererjob"
	"mapserver/web"
	"runtime"

	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/mjibson/esc -o vfs/static.go -prefix="static/" -pkg vfs static

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
		fmt.Print("OS: ")
		fmt.Println(runtime.GOOS)
		fmt.Print("Architecture: ")
		fmt.Println(runtime.GOARCH)
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
	ctx := app.Setup(p, cfg)

	//Set up mapobject events
	mapobject.Setup(ctx)

	//run initial rendering
	if ctx.Config.EnableRendering {
		go tilerendererjob.Job(ctx)
	}

	//Start http server
	//TODO: defer, may cause race condition
	web.Serve(ctx)

}
