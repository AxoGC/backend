package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestOnlines = []utils.Online{
	// 互通服务器
	{Time: time.Now().AddDate(0, 0, -1), ServerID: 1, Count: 5},
	{Time: time.Now().AddDate(0, 0, -2), ServerID: 1, Count: 12},
	{Time: time.Now().AddDate(0, 0, -3), ServerID: 1, Count: 6},
	{Time: time.Now().AddDate(0, 0, -4), ServerID: 1, Count: 8},
	{Time: time.Now().AddDate(0, 0, -5), ServerID: 1, Count: 5},
	{Time: time.Now().AddDate(0, 0, -6), ServerID: 1, Count: 8},
	{Time: time.Now().AddDate(0, 0, -7), ServerID: 1, Count: 5},

	// 基岩版服务器
	{Time: time.Now().AddDate(0, 0, -1), ServerID: 2, Count: 7},
	{Time: time.Now().AddDate(0, 0, -2), ServerID: 2, Count: 7},
	{Time: time.Now().AddDate(0, 0, -3), ServerID: 2, Count: 1},
	{Time: time.Now().AddDate(0, 0, -4), ServerID: 2, Count: 8},
	{Time: time.Now().AddDate(0, 0, -5), ServerID: 2, Count: 10},
	{Time: time.Now().AddDate(0, 0, -6), ServerID: 2, Count: 11},
	{Time: time.Now().AddDate(0, 0, -7), ServerID: 2, Count: 4},

	// 模组服务器
	{Time: time.Now().AddDate(0, 0, -1), ServerID: 3, Count: 6},
	{Time: time.Now().AddDate(0, 0, -2), ServerID: 3, Count: 10},
	{Time: time.Now().AddDate(0, 0, -3), ServerID: 3, Count: 2},
	{Time: time.Now().AddDate(0, 0, -4), ServerID: 3, Count: 3},
	{Time: time.Now().AddDate(0, 0, -5), ServerID: 3, Count: 9},
	{Time: time.Now().AddDate(0, 0, -6), ServerID: 3, Count: 8},
	{Time: time.Now().AddDate(0, 0, -7), ServerID: 3, Count: 1},

	// 饥荒联机版
	{Time: time.Now().AddDate(0, 0, -1), ServerID: 4, Count: 4},
	{Time: time.Now().AddDate(0, 0, -2), ServerID: 4, Count: 12},
	{Time: time.Now().AddDate(0, 0, -3), ServerID: 4, Count: 8},
	{Time: time.Now().AddDate(0, 0, -4), ServerID: 4, Count: 9},
	{Time: time.Now().AddDate(0, 0, -5), ServerID: 4, Count: 11},
	{Time: time.Now().AddDate(0, 0, -6), ServerID: 4, Count: 4},
	{Time: time.Now().AddDate(0, 0, -7), ServerID: 4, Count: 2},

	// 泰拉瑞亚
	{Time: time.Now().AddDate(0, 0, -1), ServerID: 5, Count: 1},
	{Time: time.Now().AddDate(0, 0, -2), ServerID: 5, Count: 5},
	{Time: time.Now().AddDate(0, 0, -3), ServerID: 5, Count: 3},
	{Time: time.Now().AddDate(0, 0, -4), ServerID: 5, Count: 4},
	{Time: time.Now().AddDate(0, 0, -5), ServerID: 5, Count: 4},
	{Time: time.Now().AddDate(0, 0, -6), ServerID: 5, Count: 2},
	{Time: time.Now().AddDate(0, 0, -7), ServerID: 5, Count: 4},
}
