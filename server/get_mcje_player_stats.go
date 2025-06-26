package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerStat struct {
	Stats       map[string]map[string]int `json:"stats"`
	DataVersion int                       `json:"DataVersion"`
}

func GetMCJEPlayerStats(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/:slug/stats/:player", []gin.HandlerFunc{
		p.Preload(cfg.Config, &p.Option{Bind: p.Uri}, nil, func(c *gin.Context, u *utils.User, r *struct {
			Slug   string `uri:"slug"`
			Player string `uri:"player"`
		}) (int, *utils.Resp) {

			var server utils.Server
			if err := cfg.DB.Take(&server, "slug = ?", r.Slug).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				return 404, Res("没有这个服务器", nil)
			} else if err != nil {
				c.Error(err)
				return 500, Res("查找服务器失败", nil)
			} else if server.Game != "mcje" {
				return 400, Res("这个服务器不是Minecraft Java版", nil)
			}

			rawUsercaches, err := os.ReadFile(filepath.Join(server.Path, "usercache.json"))
			if err != nil {
				c.Error(err)
				return 500, Res("读取玩家缓存数据失败", nil)
			}

			type UserCache struct {
				Name string `json:"name"`
				UUID string `json:"uuid"`
			}

			var userCaches []UserCache
			if err := json.Unmarshal(rawUsercaches, &userCaches); err != nil {
				c.Error(err)
				return 500, Res("解析玩家缓存数据失败", nil)
			}

			idx := slices.IndexFunc(userCaches, func(uc UserCache) bool {
				return r.Player == uc.Name
			})
			if idx == -1 {
				return 404, Res("找不到此玩家", nil)
			}

			rawStat, err := os.ReadFile(filepath.Join(server.Path, "world", "stats", userCaches[idx].UUID+".json"))
			if err != nil {
				c.Error(err)
				return 500, Res("读取玩家统计数据失败", nil)
			}

			var stat PlayerStat
			if err := json.Unmarshal(rawStat, &stat); err != nil {
				c.Error(err)
				return 500, Res("解析玩家统计数据失败", nil)
			}

			return 200, Res("", &stat)
		}),
	}
}
