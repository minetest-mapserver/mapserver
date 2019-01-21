package params

import (
	"flag"
)

type ParamsType struct {
	Help       bool
	Version    bool
	Debug      bool
	Dumpconfig bool
}

func Parse() ParamsType {
	params := ParamsType{}

	flag.BoolVar(&(params.Help), "help", false, "Show help")
	flag.BoolVar(&(params.Version), "version", false, "Show version")
	flag.BoolVar(&(params.Debug), "debug", false, "enable debug log")
	flag.Parse()

	return params
}

func PrintHelp() {
	flag.PrintDefaults()
}
