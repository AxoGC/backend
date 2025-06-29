package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
)

type HandlerConfig struct {
	*p.Config
	Env *Config
}

var Res = utils.Res
