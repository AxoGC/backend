package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
)

type HandlerConfig struct {
	*p.Config
	Env *Config
}

func HasPerm(user *utils.User, album *utils.Album) bool {

	if user.Admin {
		return true
	}

	if album.UserID == user.ID {
		return true
	}

	if album.GuildID != nil {

	}

	return false
}

type Resp = utils.Resp

var res = utils.Res
