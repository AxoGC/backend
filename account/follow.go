package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Follow(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/follow/:name", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name string `uri:"name" binding:"required"`
		}) (int, *utils.Resp) {

			if u.Name == r.Name {
				return 400, Res("不能关注自己", nil)
			}

			var userId uint
			if err := cfg.DB.Where(&utils.User{Name: r.Name}).Pluck("id", &userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("不存在这个用户", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找用户失败", nil)
			}

			if err := cfg.DB.Create(&utils.UserFollow{
				FollowerID:  u.ID,
				FollowingID: userId,
			}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
				return 409, Res("你已经关注过他了", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("关注失败", nil)
			}

			return 200, Res("关注成功", nil)
		},
	)
}
