package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddForums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/forum-groups/:forumGroupId/forums", []gin.HandlerFunc{
		p.Preload(
			&cfg.Config, &p.Option{Permission: p.Admin, Bind: p.JSON}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				ForumGroupID uint   `uri:"forumGroupId" binding:"required"`
				Slug         string `json:"slug" binding:"required"`
				Title        string `json:"title" binding:"required"`
				SubTitle     string `json:"subTitle"`
				Profile      string `json:"profile"`
				Sort         int    `json:"sort"`
			}) (int, *Resp) {

				if err := cfg.DB.Model(new(utils.Forum)).Create(r).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在相同标识或标题的论坛", nil)
				} else if errors.Is(err, gorm.ErrCheckConstraintViolated) {
					return 422, Res("不存在对应的论坛组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建论坛失败", nil)
				} else {
					return 200, Res("创建论坛成功", nil)
				}
			},
		),
	}
}
