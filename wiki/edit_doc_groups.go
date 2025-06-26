package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditDocGroups(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "PATCH", "/doc-groups/:id", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"Roles", "Roles.Role"}}, nil,
		utils.WithRolesAuth(
			[]utils.Role{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				ID    uint   `uri:"id" binding:"required"`
				Label string `json:"label" binding:"required"`
				Sort  int    `json:"sort"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Where(&utils.DocGroup{ID: r.ID}).Select("label", "sort").Updates(r); errors.Is(result.Error, gorm.ErrDuplicatedKey) {
					return 400, Res("已存在同名文档组", nil)
				} else if result.RowsAffected == 0 {
					return 400, Res("没有对应的文档组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("更新文档组失败", nil)
				} else {
					return 200, Res("更新文档组成功", nil)
				}
			},
		),
	)
}
