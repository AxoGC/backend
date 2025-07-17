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
		UserID:         1,
		ReviewableID:   1,
		ReviewableType: "Doc",
		ReviewCount:    2,
	},
	{
		ID:             2,
		UpdatedAt:      time.Now().AddDate(0, 0, -3),
		Content:        "红石电路确实有点复杂，不过讲解得很清楚",
		Attitude:       boolPtr(true),
		UserID:         1,
		ReviewableID:   2,
		ReviewableType: "docs",
		ReviewCount:    0,
	},
	{
		ID:             3,
		UpdatedAt:      time.Now().AddDate(0, 0, -2),
		Content:        "城堡建筑太震撼了！求教程！",
		Attitude:       boolPtr(true),
		UserID:         1,
		ReviewableID:   1,
		ReviewableType: "albums",
		ReviewCount:    1,
	},
	{
		ID:             4,
		UpdatedAt:      time.Now().AddDate(0, 0, -1),
		Content:        "Boss攻略很实用，帮助我通关了好几个boss",
		Attitude:       boolPtr(true),
		UserID:         1,
		ReviewableID:   4,
		ReviewableType: "posts",
		ReviewCount:    0,
	},
}
