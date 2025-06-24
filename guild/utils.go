package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
)

const (
	None      utils.GuildRole = "none"
	Applicant utils.GuildRole = "applicant"
	Member    utils.GuildRole = "member"
	Admin     utils.GuildRole = "admin"
	Owner     utils.GuildRole = "owner"
)

type HandlerConfig struct {
	*p.Config
	Env *Config
}

type Resp = utils.Resp

var Res = utils.Res
