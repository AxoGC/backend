package main

import (
	"strconv"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func SetCover(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/cover", p.Preload(
		cfg.Config, &p.Option{Login: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {
			return utils.UploadImageMidWare(cfg.Env.FilePath, "user-cover/", strconv.Itoa(int(u.ID)))(c)
		},
	)
}
