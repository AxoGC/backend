package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
)

func RegisterCheckin(r *gin.Engine, cfg *p.Config) {

	r.GET("/checkin", p.Preload(
		cfg, &p.Option{Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) {

		},
	))

	r.POST("/checkin", p.Preload(
		cfg, &p.Option{Permission: p.Login}, nil,
		func(c *gin.Context, u *utils.User, r *struct{}) {

		},
	))
}
