package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Kick(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/kick", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, Admin, Owner), p.Preload(
			cfg.Config, &p.Option{}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				IDs []uint `json:"ids"`
			}) (int, *Resp) {

				if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

					result := tx.Where(
						"id IN ? AND guild_id = ? AND guild_role = ?", r.IDs, u.GuildID, Member,
					).Updates(map[string]any{"guild_id": nil, "guild_role": None})

					if result.Error != nil {
						c.Error(result.Error)
						return utils.TxRes(500, Res("踢出成员失败", nil))
					}

					if err := tx.Where(&utils.Guild{ID: *u.GuildID}).Update(
						"count", gorm.Expr("count - ?", result.RowsAffected),
					).Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("修改公会成员失败", nil))
					}

					return nil
				}).(*utils.TxResp[Resp]); ok {
					return err.Code, Res(err.Data.Message, nil)
				}

				return 200, Res("成员踢出成功", nil)
			},
		),
	}
}
