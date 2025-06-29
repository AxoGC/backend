package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestUserGuilds = []utils.UserGuild{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -4, 0),
		UserID:    2, // 史蒂夫
		Status:    boolPtr(true),
		Admin:     true,
		GuildID:   1, // 建筑师联盟
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, -3, 0),
		UserID:    3, // 艾莉克斯
		Status:    boolPtr(true),
		Admin:     false,
		GuildID:   3, // 冒险者工会
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		UserID:    4, // 农夫威尔逊
		Status:    boolPtr(true),
		Admin:     false,
		GuildID:   3, // 冒险者工会
	},
}
