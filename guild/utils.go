package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
)

const (
	Applicant = utils.GuildApplicant
	Member    = utils.GuildMember
	Admin     = utils.GuildAdmin
)

type HandlerConfig struct {
	*p.Config
	Env *Config
}

var Res = utils.Res
