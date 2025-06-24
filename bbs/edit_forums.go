package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditForums(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "PATCH", "/forums/:slug", []gin.HandlerFunc{p.Preload(
		&cfg.Config, &p.Option{Permission: p.Admin, Bind: p.Uri | p.JSON}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Uri          string `uri:"slug" binding:"required,min=3,alphanum"`
			Slug         string `json:"slug" binding:"required,min=3,alphanum"`
			ForumGroupID uint   `json:"forumGroupId" binding:"required"`
			Title        string `json:"title"`
			SubTitle     string `json:"subTitle"`
			Profile      string `json:"profile"`
			Sort         int    `json:"sort"`
		}) (int, *Resp) {

			if result := cfg.DB.Where(&utils.Forum{Slug: r.Uri}).Updates(r); errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
				return 404, Res("不存在对应的论坛组", nil)
			} else if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
				return 409, Res("该标题或标识已被其他论坛使用", nil)
			} else if result.RowsAffected == 0 {
				return 404, Res("不存在该论坛", nil)
			} else if result.Error != nil {
				c.Error(result.Error)
				return 500, Res("论坛修改失败", nil)
			} else {
				return 200, Res("论坛修改成功", nil)
			}
		},
	),
	}
}
