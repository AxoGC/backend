package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddImages(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "POST", "/images", []gin.HandlerFunc{
		p.Preload(
			cfg.Config, &p.Option{Permission: p.Login, Bind: p.Other}, nil,
			func(c *gin.Context, u *utils.User, r *struct {
				AlbumID uint   `form:"albumId"`
				Label   string `form:"label"`
				Rename  bool   `form:"rename"`
			}) (int, *Resp) {

				var album utils.Album
				if err := cfg.DB.Take(&album, r.AlbumID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(404, Resp{"找不到对应的相册", nil})
					return
				} else if err != nil {
					c.JSON(500, Resp{"查找相册失败", nil})
					c.Error(err)
					return
				}

				if !u.Admin {
					c.JSON(403, Resp{"权限不足", nil})
					return
				}
			},
		),
	}
}

func CanUpload(user utils.User, album utils.Album, isGuildAdmin bool) bool {

	if user.Admin || user.ID == album.UserID {
		return true
	}

	if album.GuildID != nil {
		if user.GuildID == nil || *user.GuildID != *album.GuildID {
			return false
		}

		if album.Protected {
			return isGuildAdmin
		} else {
			return true
		}
	}

	if album.Protected {
		return false
	}

	return true
}
