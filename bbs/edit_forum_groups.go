package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditForumGroups(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "PUT", "/forum-groups/:id", p.Preload(
		&cfg.Config, &p.Option{Login: p.Login, Bind: p.URI | p.JSON}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.BBSAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				ID    uint   `uri:"id"`
				Label string `json:"label"`
				Sort  int    `json:"sort"`
			}) (int, *utils.Resp) {

				if result := cfg.DB.Model(new(utils.ForumGroup)).Where(r.ID).Updates(map[string]any{
					"label": r.Label,
					"sort":  r.Sort,
				}); errors.Is(result.Error, gorm.ErrDuplicatedKey) {
					return 409, Res("此标题已被其他论坛组使用", nil)
				} else if result.RowsAffected == 0 {
					return 404, Res("不存在此论坛组", nil)
				} else if result.Error != nil {
					c.Error(result.Error)
					return 500, Res("更新论坛组失败", nil)
				} else {
					return 200, Res("更新论坛组成功", nil)
				}
			},
		),
	)
}
