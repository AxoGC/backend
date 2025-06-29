package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/forum-groups", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				Label string `json:"label"`
				Sort  int    `json:"sort"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Model(new(utils.ForumGroup)).Create(map[string]any{
					"label": r.Label,
					"sort":  r.Sort,
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在同名论坛组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("论坛组创建失败", nil)
				} else {
					return 201, Res("论坛组创建成功", nil)
				}
			},
		),
	)
}
