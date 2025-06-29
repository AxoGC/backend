package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/redis/go-redis/v9"
)

type HandlerConfig struct {
	*p.Config
	Env *Config
	RDB *redis.Client
	BCs map[string][]BedrockCommand
	BRs map[string]chan BedrockResponse
}

var Res = utils.Res
