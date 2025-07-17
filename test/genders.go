package main

import "github.com/axogc/backend/utils"

var TestGenders = []utils.Gender{
	{
		ID:    utils.UnknownGender,
		Label: "未知性别",
	},
	{
		ID:    utils.Male,
		Label: "男",
	},
	{
		ID:    utils.Female,
		Label: "女",
	},
	{
		ID:    utils.Femboy,
		Label: "小男娘",
	},
}
