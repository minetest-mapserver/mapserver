package params

import (
	"flag"
	"fmt"
)

type ParamsType struct {
	worlddir string
}

var params ParamsType

func Params() ParamsType {
	return params
}

func Parse(){
	params := ParamsType{}
	flag.StringVar(&(params.worlddir), "worlddir", "./", "world directory")
	flag.Parse()
	fmt.Println("World dir is: ", params.worlddir)
}