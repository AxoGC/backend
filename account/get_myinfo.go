package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
)

func GetMyinfo(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/myinfo", p.Preload(
		cfg.Config, &p.Option{Login: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {
			return 200, Res("", gin.H{
				"id":   u.ID,
				"name": u.Name,
			})
		},
	)
}
