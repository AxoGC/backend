package main

import (
	"time"

	"github.com/axogc/backend/utils"
	"github.com/samber/lo"
)

var TestLogs = []utils.Log{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, 0, 0),
		Path:      "/api/users/login",
		Method:    "POST",
		Status:    200,
		UserID:    lo.ToPtr(uint(1)),
		Message:   "用户登录成功",
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, 0, -1),
		Path:      "/api/docs/create",
		Method:    "POST",
		Status:    201,
		UserID:    lo.ToPtr(uint(1)),
		Message:   "创建文档成功",
	},
}
