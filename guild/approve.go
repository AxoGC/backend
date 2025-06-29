package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Approve(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/approve", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.JSON}, nil,
		WithGuildRolesAuth([]utils.DictID{Admin, Owner},
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids"`
			}) (int, *utils.Resp) {

				if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

					result := tx.Where(
						"id IN ? AND guild_id IN ? AND guild_role = ?", r.IDs, u.GuildID, Applicant,
					).Update("guild_role", Member)

					if result.Error != nil {
						c.Error(result.Error)
						return utils.TxRes(500, Res("批准申请失败", nil))
					}

					if err := tx.Where(&utils.Guild{ID: *u.GuildID}).Update(
						"count", gorm.Expr("count - ?", result.RowsAffected),
					).Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("修改公会成员失败", nil))
					}

					return nil
				}).(*utils.TxResp[utils.Resp]); ok {
					return err.Code, Res(err.Data.Message, nil)
				}

				return 200, Res("批准申请成功", nil)
			},
		),
	)
}
