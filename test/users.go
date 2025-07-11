package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestUsers = []utils.User{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -6, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -1),
		Name:      "管理员",
		Exp:       9999,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye", // 示例加密密码
		Gender:    boolPtr(true),
		Profile:   "系统管理员，负责维护整个社区的正常运行",
		Birthday:  timePtr(time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)),
		Location:  "北京市",
		DailyCoin: 1000,
		HonorCoin: 5000,
		Checkin:   365,
		Email:     "admin@gameclub.com",
		QQ:        stringPtr("123456789"),
		MCBEName:  stringPtr("AdminMCBE"),
		MCJEName:  stringPtr("AdminMCJE"),
		Setting: map[string]any{
			"theme":        "dark",
			"language":     "zh-CN",
			"notification": true,
		},
		FollowingCount: 10,
		FollowerCount:  50,
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, -3, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -2),
		Name:      "史蒂夫",
		Exp:       2580,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		Gender:    boolPtr(true),
		Profile:   "我的世界资深玩家，喜欢建造和红石科技",
		Birthday:  timePtr(time.Date(1995, 8, 20, 0, 0, 0, 0, time.UTC)),
		Location:  "上海市",
		DailyCoin: 520,
		HonorCoin: 1200,
		Checkin:   90,
		Email:     "steve@gameclub.com",
		QQ:        stringPtr("987654321"),
		MCBEName:  stringPtr("Steve_Builder"),
		MCJEName:  stringPtr("SteveRedstone"),
		Setting: map[string]any{
			"theme":        "light",
			"language":     "zh-CN",
			"notification": true,
		},
		FollowingCount: 25,
		FollowerCount:  30,
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -3),
		Name:      "艾莉克斯",
		Exp:       1800,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		Gender:    boolPtr(false),
		Profile:   "多游戏爱好者，特别喜欢冒险和探索类游戏",
		Birthday:  timePtr(time.Date(1998, 12, 3, 0, 0, 0, 0, time.UTC)),
		Location:  "广州市",
		DailyCoin: 280,
		HonorCoin: 800,
		Checkin:   45,
		Email:     "alex@gameclub.com",
		MCBEName:  stringPtr("Alex_Explorer"),
		Setting: map[string]any{
			"theme":        "auto",
			"language":     "zh-CN",
			"notification": false,
		},
		FollowingCount: 15,
		FollowerCount:  20,
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		UpdatedAt: time.Now(),
		Name:      "农夫威尔逊",
		Exp:       950,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		Gender:    boolPtr(true),
		Profile:   "星露谷物语和饥荒的忠实粉丝",
		Birthday:  timePtr(time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC)),
		Location:  "成都市",
		DailyCoin: 150,
		HonorCoin: 300,
		Checkin:   30,
		Email:     "wilson@gameclub.com",
		QQ:        stringPtr("555666777"),
		Setting: map[string]any{
			"theme":        "light",
			"language":     "zh-CN",
			"notification": true,
		},
		FollowingCount: 8,
		FollowerCount:  12,
	},
	{
		ID:        5,
		CreatedAt: time.Now().AddDate(0, 0, -15),
		UpdatedAt: time.Now().AddDate(0, 0, -1),
		Name:      "恐龙猎人",
		Exp:       650,
		Password:  "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		Gender:    boolPtr(true),
		Profile:   "方舟生存进化专家，恐龙驯服大师",
		Location:  "杭州市",
		DailyCoin: 80,
		HonorCoin: 150,
		Checkin:   15,
		Email:     "dinohunter@gameclub.com",
		Setting: map[string]any{
			"theme":        "dark",
			"language":     "zh-CN",
			"notification": true,
		},
		FollowingCount: 5,
		FollowerCount:  8,
	},
}
