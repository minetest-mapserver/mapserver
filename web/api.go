package web

import "mapserver/app"

type Api struct {
	Context *app.App
}

func NewApi(ctx *app.App) *Api {
	return &Api{
		Context: ctx,
	}
}
