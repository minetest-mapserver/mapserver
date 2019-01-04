package main

import (
	"mapserver/params"
)

func main(){
	params.Parse()
	p := params.Params()
	if (p.Help){
		return
	}
}