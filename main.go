package main

import (
	"mapserver/params"
	"mapserver/worldconfig"
	"flag"
	"fmt"
)

var (
	Version string
)

func main() {
	p := params.Parse()

	if p.Help {
		flag.PrintDefaults()
		return
	}

	if p.Version {
		fmt.Print("Mapserver version: ")
		if Version == "" {
			Version = "SNAPSHOT"
		}
		fmt.Println(Version)
		return
	}

	worldcfg := worldconfig.Parse(p.Worlddir + "world.mt")
	fmt.Println("Config ", worldcfg)
}
