package params

import (
	"flag"
)

type ParamsType struct {
	Help         bool
	Version      bool
	Debug        bool
	CreateConfig bool
}

func Parse() ParamsType {
	params := ParamsType{}

	flag.BoolVar(&(params.Help), "help", false, "Show help")
	flag.BoolVar(&(params.Version), "version", false, "Show version")
	flag.BoolVar(&(params.Debug), "debug", false, "enable debug log")
	flag.BoolVar(&(params.CreateConfig), "createconfig", false, "creates a config and exits")
	flag.Parse()

	return params
}

func PrintHelp() {
	flag.PrintDefaults()
}
