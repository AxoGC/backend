package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestAlbums = []utils.Album{
	{
		ID:          1,
		CreatedAt:   time.Now().AddDate(0, -2, 0),
		UpdatedAt:   time.Now().AddDate(0, 0, -5),
		UserID:      2,          // 史蒂夫
		GuildID:     uintPtr(1), // 建筑师联盟
		Slug:        "medieval-castle",
		Label:       "中世纪城堡建筑",
		Profile:     "我花了一个月时间建造的大型中世纪城堡，包含完整的内部装饰",
		Pinned:      true,
		Hide:        false,
		Protected:   false,
		ImageCount:  12,
		ReviewCount: 18,
	},
	{
		ID:          2,
		CreatedAt:   time.Now().AddDate(0, -1, 0),
		UpdatedAt:   time.Now().AddDate(0, 0, -3),
		UserID:      3,          // 艾莉克斯
		GuildID:     uintPtr(3), // 冒险者工会
		Slug:        "adventure-screenshots",
		Label:       "冒险截图集",
		Profile:     "各种游戏的精彩冒险瞬间记录",
		Pinned:      false,
		Hide:        false,
		Protected:   false,
		ImageCount:  25,
		ReviewCount: 9,
	},
	{
		ID:          3,
		CreatedAt:   time.Now().AddDate(0, 0, -15),
		UpdatedAt:   time.Now().AddDate(0, 0, -1),
		UserID:      4, // 威尔逊
		GuildID:     nil,
		Slug:        "farm-designs",
		Label:       "农场设计图",
		Profile:     "星露谷物语和我的世界的各种农场设计方案",
		Pinned:      false,
		Hide:        false,
		Protected:   true,
		ImageCount:  8,
		ReviewCount: 6,
	},
}
