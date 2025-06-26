package main

import p "github.com/bestcb2333/gin-gorm-preloader/v2"

type HandlerConfig struct {
	*p.Config
	Env *Config
}

type Resp struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
