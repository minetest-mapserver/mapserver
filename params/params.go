package params

import (
	"flag"
)

type ParamsType struct {
	Worlddir string
	Port     int
	Help     bool
	Version  bool
}

func Parse() ParamsType {
	params := ParamsType{}

	flag.StringVar(&(params.Worlddir), "worlddir", "./", "world directory")
	flag.IntVar(&(params.Port), "port", 8080, "port to use")
	flag.BoolVar(&(params.Help), "help", false, "Show help")
	flag.BoolVar(&(params.Version), "version", false, "Show version")
	flag.Parse()

	return params
}

func PrintHelp(){
	flag.PrintDefaults()
}
