package main

import (
	"errors"
	"fmt"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transfer(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/transfer", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON}, nil,
		WithGuildRolesAuth(
			[]utils.DictID{Owner},
			func(c *gin.Context, u *utils.User, r *struct {
				ID uint `json:"id"`
			}) (int, *utils.Resp) {

				var user utils.User
				if err := cfg.DB.Take(&user, r.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					return 404, Res("找不到这个用户", nil)
				} else if err != nil {
					c.Error(err)
					return 500, Res("查询用户失败", nil)
				}

				if user.GuildID != u.GuildID {
					return 400, Res(fmt.Sprintf("用户 %s 不属于此公会", user.Name), nil)
				}

				if user.GuildRoleID != Member && user.GuildRoleID != Admin {
					return 400, Res(fmt.Sprintf("用户 %s 角色 %s 不满足转让条件", user.Name, user.GuildRole), nil)
				}

				if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

					if err := tx.Save(&utils.User{ID: u.ID, GuildRoleID: Admin}).Error; err != nil {
						return utils.TxRes(500, Res("我的公会角色降低失败", nil))
					}

					if err := tx.Save(&utils.User{ID: user.ID, GuildRoleID: Owner}).Error; err != nil {
						return utils.TxRes(500, Res("目标公会角色降低失败", nil))
					}

					return nil
				}).(*utils.TxResp[utils.Resp]); ok {
					return err.Code, Res(err.Data.Message, nil)
				}

				return 200, Res("公会转让成功", nil)
			},
		),
	)
}
