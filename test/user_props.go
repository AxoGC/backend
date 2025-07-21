package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestUserProps = []utils.UserProp{
	{
		ID:        1,
		UpdatedAt: time.Now().AddDate(0, 0, -10),
		PropID:    utils.MysteryBox,
		UserID:    1, // 管理员
		Count:     10,
	},
	{
		ID:        2,
		UpdatedAt: time.Now().AddDate(0, 0, -5),
		PropID:    utils.MysteryBox,
		UserID:    1, // 史蒂夫
		Count:     3,
	},
	{
		ID:        3,
		UpdatedAt: time.Now().AddDate(0, 0, -3),
		PropID:    utils.MysteryBox,
		UserID:    1, // 艾莉克斯
		Count:     5,
	},
}
