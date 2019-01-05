package main

import (
	"mapserver/params"
	"mapserver/worldconfig"
	"fmt"
)

func main() {
	p := params.Parse()
	if p.Help {
		return
	}

	worldcfg := worldconfig.Parse(p.Worlddir + "world.mt")
	fmt.Println("Config ", worldcfg)
}
