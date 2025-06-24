package main

import (
	"strconv"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetAvatar(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/avatar", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login}, nil,
			func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {
				return utils.UploadImageMidWare(cfg.Env.FilePath, "user-avatar/", strconv.Itoa(int(u.ID)))(c)
			},
		),
	}
}
