package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestOnlines = []utils.Online{
	{
		ID:       1,
		Time:     time.Now().AddDate(0, 0, 0).Add(-1 * time.Hour),
		ServerID: 1, // 生存主世界
		Count:    35,
	},
	{
		ID:       2,
		Time:     time.Now().AddDate(0, 0, 0).Add(-2 * time.Hour),
		ServerID: 1,
		Count:    28,
	},
	{
		ID:       3,
		Time:     time.Now().AddDate(0, 0, 0).Add(-3 * time.Hour),
		ServerID: 2, // 创造建筑服
		Count:    15,
	},
	{
		ID:       4,
		Time:     time.Now().AddDate(0, 0, 0).Add(-1 * time.Hour),
		ServerID: 3, // 手机版联机
		Count:    12,
	},
	{
		ID:       5,
		Time:     time.Now().AddDate(0, 0, 0).Add(-4 * time.Hour),
		ServerID: 4, // 泰拉瑞亚专家
		Count:    6,
	},
}
