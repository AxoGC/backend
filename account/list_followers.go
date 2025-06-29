package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListFollowers(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/users/:name/followers", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name string `uri:"name" binding:"required"`
		}) (int, *utils.Resp) {

			var users []utils.UserPreview
			if err := cfg.DB.Model(new(utils.User)).Preload("Followers", s.Model(new(utils.User))).Take(&struct {
				Followers *[]utils.UserPreview `json:"followers" gorm:"many2many:user_follows;joinForeignKey:following_id;joinReferences:follower_id"`
			}{&users}, "name = ?", r.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到对应的用户", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找用户失败", nil)
			}

			return 200, Res("获取粉丝列表成功", users)
		},
	)
}
