package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestLogs = []utils.Log{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, 0, 0),
		Path:      "/api/users/login",
		Method:    "POST",
		Status:    200,
		UserID:    uintPtr(2),
		Message:   "用户登录成功",
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, 0, -1),
		Path:      "/api/docs/create",
		Method:    "POST",
		Status:    201,
		UserID:    uintPtr(2),
		Message:   "创建文档成功",
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, 0, -1),
		Path:      "/api/upload",
		Method:    "POST",
		Status:    413,
		UserID:    uintPtr(3),
		Message:   "文件大小超出限制",
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, 0, -2),
		Path:      "/api/users/register",
		Method:    "POST",
		Status:    400,
		UserID:    nil,
		Message:   "邮箱格式不正确",
	},
}
