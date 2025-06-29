package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestDonations = []utils.Donation{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		Amount:    99.99,
		UserID:    uintPtr(2), // 史蒂夫
		Message:   "感谢提供这么棒的游戏社区！",
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, 0, -15),
		Amount:    66.66,
		UserID:    uintPtr(3), // 艾莉克斯
		Message:   "支持社区发展，希望越来越好！",
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, 0, -30),
		Amount:    188.88,
		UserID:    nil, // 匿名捐赠
		Message:   "匿名支持，加油！",
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, 0, -7),
		Amount:    50.00,
		UserID:    uintPtr(4), // 威尔逊
		Message:   "小小心意，不成敬意",
	},
}
