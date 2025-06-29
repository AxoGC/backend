package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestReviews = []utils.Review{
	{
		ID:             1,
		UpdatedAt:      time.Now().AddDate(0, 0, -5),
		Content:        "这个教程写得非常详细，对新手很友好！",
		Attitude:       boolPtr(true),
		UserID:         3, // 艾莉克斯
		ReviewableID:   1, // 对文档1的评论
		ReviewableType: "Doc",
		ReviewCount:    2,
	},
	{
		ID:             2,
		UpdatedAt:      time.Now().AddDate(0, 0, -3),
		Content:        "红石电路确实有点复杂，不过讲解得很清楚",
		Attitude:       boolPtr(true),
		UserID:         4, // 威尔逊
		ReviewableID:   2, // 对文档2的评论
		ReviewableType: "docs",
		ReviewCount:    0,
	},
	{
		ID:             3,
		UpdatedAt:      time.Now().AddDate(0, 0, -2),
		Content:        "城堡建筑太震撼了！求教程！",
		Attitude:       boolPtr(true),
		UserID:         5, // 恐龙猎人
		ReviewableID:   1, // 对相册1的评论
		ReviewableType: "albums",
		ReviewCount:    1,
	},
	{
		ID:             4,
		UpdatedAt:      time.Now().AddDate(0, 0, -1),
		Content:        "Boss攻略很实用，帮助我通关了好几个boss",
		Attitude:       boolPtr(true),
		UserID:         2, // 史蒂夫
		ReviewableID:   4, // 对文档4的评论
		ReviewableType: "docs",
		ReviewCount:    0,
	},
}
