package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestImages = []utils.Image{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		Filename:  "castle_main_view_001.jpg",
		Label:     "城堡主视图",
		Likes:     45,
		UserID:    2, // 史蒂夫
		AlbumID:   1, // 中世纪城堡建筑
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		Filename:  "castle_interior_001.jpg",
		Label:     "城堡内部大厅",
		Likes:     32,
		UserID:    2,
		AlbumID:   1,
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		Filename:  "castle_tower_001.jpg",
		Label:     "城堡塔楼",
		Likes:     28,
		UserID:    2,
		AlbumID:   1,
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		Filename:  "terraria_boss_fight_001.jpg",
		Label:     "泰拉瑞亚Boss战",
		Likes:     18,
		UserID:    3, // 艾莉克斯
		AlbumID:   2, // 冒险截图集
	},
	{
		ID:        5,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		Filename:  "stardew_farm_001.jpg",
		Label:     "星露谷农场",
		Likes:     22,
		UserID:    3,
		AlbumID:   2,
	},
	{
		ID:        6,
		CreatedAt: time.Now().AddDate(0, 0, -15),
		Filename:  "auto_farm_design_001.jpg",
		Label:     "自动农场设计图",
		Likes:     15,
		UserID:    4, // 威尔逊
		AlbumID:   3, // 农场设计图
	},
}
