package main

import (
	"mapserver/params"
	"mapserver/worldconfig"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	Version = "2.0-DEV"
)

func main() {
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
	logrus.WithFields(logrus.Fields{"version":Version}).Info("Starting mapserver")

	fmt.Println("Config ", worldcfg)
}
