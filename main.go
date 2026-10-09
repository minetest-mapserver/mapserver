package main

import (
	"fmt"
	"mapserver/app"
	"mapserver/mapobject"
	"mapserver/tilerendererjob"
	"mapserver/util"
	"mapserver/web"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func main() {
	//get Env or Default value
	env := func(key string, defaultVal string) string {
		val, ok := os.LookupEnv(key)
		if !ok {
			return defaultVal
		}
		return val
	}

	//Parse command line
	p := util.ParseParams()

	if p.Debug || env("MT_LOGLEVEL", "INFO") == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if p.Help {
		util.PrintHelp()
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
	cfg, err := app.ParseConfig(env("MT_CONFIG_PATH", app.ConfigFile))
	if err != nil {
		panic(err)
	}

	//write back config with all values
	if env("MT_READONLY", "false") != "true" {
		err = cfg.Save()
		if err != nil {
			panic(err)
		}
	}

	//exit after creating the config
	if p.CreateConfig {
		return
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
	web.Serve(ctx)

}
