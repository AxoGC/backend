package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Unfollow(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "DELETE", "/unfollow/:name", p.Preload(
		cfg.Config, &p.Option{Login: p.Login, Bind: p.URI}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Name string `uri:"name" binding:"required"`
		}) (int, *utils.Resp) {

			var userId uint
			if err := cfg.DB.Where(&utils.User{Name: r.Name}).Pluck("id", &userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("找不到该用户", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找该用户失败", nil)
			}

			if result := cfg.DB.Delete(new(utils.UserFollow), "follower_id = ? AND following_id = ?", u.ID, userId); result.RowsAffected == 0 {
				return 400, Res("你还没有关注该用户", nil)
			} else if result.Error != nil {
				c.Error(result.Error)
				return 500, Res("取消关注失败", nil)
			}

			return 200, Res("取消关注成功", nil)
		},
	)
}
