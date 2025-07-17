package main

import (
	"github.com/axogc/backend/utils"
)

var TestDonations = []utils.Donation{
	{
		ID:      1,
		Amount:  99.99,
		UserID:  uintPtr(1),
		Message: "感谢提供这么棒的游戏社区！",
	},
	{
		ID:      2,
		Amount:  66.66,
		UserID:  nil,
		Message: "支持社区发展，希望越来越好！",
	},
	{
		ID:      3,
		Amount:  188.88,
		UserID:  nil,
		Message: "匿名支持，加油！",
	},
	{
		ID:      4,
		Amount:  50.00,
		UserID:  nil,
		Message: "小小心意，不成敬意",
	},
}
