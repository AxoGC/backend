package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDocs(cfg *p.Config) (string, string, gin.HandlerFunc) {
	return "POST", "/doc-groups/:docGroupId/docs", p.Preload(
		cfg, &p.Option{Login: p.Login, Bind: p.URI | p.JSON, Preloads: []string{"UserRoles"}}, nil,
		utils.WithRolesAuth(
			[]utils.RoleID{utils.Admin, utils.WikiAdmin},
			func(c *gin.Context, u *utils.User, r *struct {
				DocGroupID uint   `uri:"docGroupId"`
				Slug       string `json:"slug"`
				Title      string `json:"title"`
				Content    string `json:"content"`
				Sort       int    `json:"sort"`
			}) (int, *utils.Resp) {

				if err := cfg.DB.Model(new(utils.Doc)).Create(map[string]any{
					"doc_group_id": r.DocGroupID,
					"slug":         r.Slug,
					"title":        r.Title,
					"content":      r.Content,
					"sort":         r.Sort,
					"user_id":      u.ID,
				}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
					return 409, Res("已存在相同名称或标识的文档", nil)
				} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return 404, Res("不存在对应的文档组", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("创建文档失败", nil)
				} else {
					return 201, Res("文档创建成功", nil)
				}
			},
		),
	)
}
