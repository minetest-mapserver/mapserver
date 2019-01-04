package params

import (
	"flag"
)

type ParamsType struct {
	Worlddir string
	Port     int
	Help     bool
}

var params ParamsType

func Params() ParamsType {
	return params
}

func Parse() {
	params := ParamsType{}

	flag.StringVar(&(params.Worlddir), "worlddir", "./", "world directory")
	flag.IntVar(&(params.Port), "port", 8080, "port to use")
	flag.BoolVar(&(params.Help), "help", false, "Show help")
	flag.Parse()

	if params.Help {
		flag.PrintDefaults()
	}
}
