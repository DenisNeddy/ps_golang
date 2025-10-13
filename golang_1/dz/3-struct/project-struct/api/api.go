package api

import "project/download-json/config"

type Api struct {
	config string
}

func NewApi() *Api {
	conf := config.NewConfig()
	return &Api{
		config: conf.Key,
	}
}
