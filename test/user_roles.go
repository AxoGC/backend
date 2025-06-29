package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestUserRoles = []utils.UserRole{
	{
		ID:        1,
		CreatedAt: time.Now().AddDate(0, -6, 0),
		UserID:    1, // 管理员拥有所有权限
		RoleID:    utils.Admin,
	},
	{
		ID:        2,
		CreatedAt: time.Now().AddDate(0, -3, 0),
		UserID:    2, // 史蒂夫是知识库管理员
		RoleID:    utils.WikiAdmin,
	},
	{
		ID:        3,
		CreatedAt: time.Now().AddDate(0, -2, 0),
		UserID:    3, // 艾莉克斯是论坛管理员
		RoleID:    utils.BBSAdmin,
	},
	{
		ID:        4,
		CreatedAt: time.Now().AddDate(0, -1, 0),
		UserID:    2, // 史蒂夫也是图库管理员
		RoleID:    utils.GalleryAdmin,
	},
}
