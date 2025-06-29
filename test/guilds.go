package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestGuilds = []utils.Guild{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -4, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -5),
		Name:      "建筑师联盟",
		Slug:      "builders-union",
		UserCount: 25,
		Profile:   "专注于游戏建筑设计和创作的玩家公会",
		Notice:    "欢迎所有热爱建筑的玩家加入我们！每周日晚8点有建筑比赛活动。",
		Money:     50000,
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -3),
		Name:      "红石科技社",
		Slug:      "redstone-tech",
		UserCount: 18,
		Profile:   "致力于研究红石电路和自动化装置的技术公会",
		Notice:    "新成员请先阅读红石基础教程，技术交流群：123456789",
		Money:     32000,
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		UpdatedAt: time.Now().AddDate(0, 0, -1),
		Name:      "冒险者工会",
		Slug:      "adventurers-guild",
		UserCount: 40,
		Profile:   "多游戏冒险团队，组织各种游戏的团队活动",
		Notice:    "本周六下午2点组织泰拉瑞亚boss挑战，有兴趣的成员请报名！",
		Money:     75000,
	},
}
