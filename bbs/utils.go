package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
)

type HandlerConfig struct {
	p.Config
	Env *Config
}

type Resp = utils.Resp

var Res = utils.Res
