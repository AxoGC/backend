package main

import (
	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Quit(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/quit", []gin.HandlerFunc{
		CheckRoleMidWare(cfg.Config, Member, Admin), p.Preload(
			cfg.Config, &p.Option{Permission: p.Login}, nil,
			func(c *gin.Context, u *utils.User, r *struct{}) (int, *Resp) {

				if err, ok := cfg.DB.Transaction(func(tx *gorm.DB) error {

					if err := tx.Where(&utils.Guild{ID: *u.GuildID}).Update("count", "count - 1").Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("修改公会人数失败", nil))
					}

					if err := tx.Where(u).Updates(
						map[string]any{"guild_id": nil, "guild_role": 0},
					).Error; err != nil {
						c.Error(err)
						return utils.TxRes(500, Res("退出公会失败", nil))
					}

					return nil
				}).(*utils.TxResp[Resp]); ok {
					return err.Code, Res(err.Data.Message, nil)
				}

				return 200, Res("退出公会成功", nil)
			},
		),
	}
}
