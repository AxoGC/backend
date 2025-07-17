package main

import (
	"errors"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	s "github.com/bestcb2333/scoper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "GET", "/albums/:slug", p.Preload(
		cfg.Config, &p.Option{Bind: p.URI | p.Query, Login: p.Auto}, nil,
		func(c *gin.Context, u *utils.User, r *struct {
			Slug string `uri:"slug"`
			All  bool   `form:"all"`
		}) (int, *utils.Resp) {

			type User struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}

			type Image struct {
				ID       uint   `json:"id"`
				Filename string `json:"filename"`
				Label    string `json:"label"`
				Likes    uint   `json:"likes"`
				AlbumID  uint   `json:"albumId"`
			}

			type Review struct {
				ID             uint   `json:"id"`
				Content        string `json:"content"`
				UserID         uint   `json:"userId"`
				ReviewableID   uint   `json:"reviewableId"`
				ReviewableType string `json:"reviewableType"`
				User           User   `json:"user"`
			}

			type Album struct {
				ID          uint     `json:"id"`
				Label       string   `json:"label"`
				Slug        string   `json:"slug"`
				Profile     string   `json:"profile"`
				UserID      uint     `json:"userId"`
				User        User     `json:"user"`
				ReviewCount uint     `json:"reviewCount"`
				Images      []Image  `json:"images"`
				Reviews     []Review `json:"reviews" gorm:"polymorphic:Reviewable;polymorphicValue:albums"`
			}

			var album Album

			query := cfg.DB.Model(new(utils.Album)).Preload("User",
				s.Model(new(utils.User)),
			).Preload("Images",
				s.Model(new(utils.Image)),
				utils.Paginate(c, nil),
			).Preload("Reviews",
				s.Model(new(utils.Review)),
				utils.Paginate(c, &utils.PaginateConfig{KeyPrefix: "review"}),
				s.Preload("User"),
			)

			if !r.All || !(u != nil && u.HasAnyRole(utils.Admin, utils.GalleryAdmin)) {
				query = query.Where("hide = ?", false)
			}

			if err := query.First(&album, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 400, Res("不存在这个相册", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("获取相册失败", nil)
			}

			return 200, Res("", &album)
		},
	)
}
