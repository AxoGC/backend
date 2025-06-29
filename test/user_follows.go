package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestUserFollows = []utils.UserFollow{
	{
		ID:          1,
		CreatedAt:   time.Now().AddDate(0, -2, 0),
		FollowerID:  3, // 艾莉克斯关注史蒂夫
		FollowingID: 2,
	},
	{
		ID:          2,
		CreatedAt:   time.Now().AddDate(0, -1, -15),
		FollowerID:  4, // 威尔逊关注史蒂夫
		FollowingID: 2,
	},
	{
		ID:          3,
		CreatedAt:   time.Now().AddDate(0, -1, -10),
		FollowerID:  5, // 恐龙猎人关注艾莉克斯
		FollowingID: 3,
	},
	{
		ID:          4,
		CreatedAt:   time.Now().AddDate(0, -1, -5),
		FollowerID:  2, // 史蒂夫关注管理员
		FollowingID: 1,
	},
}
