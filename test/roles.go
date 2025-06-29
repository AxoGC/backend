package main

import "github.com/axogc/backend/utils"

var TestRoles = []utils.Role{
	{
		ID:          utils.Admin,
		Label:       "超级管理员",
		Description: "拥有系统所有权限的超级管理员",
	},
	{
		ID:          utils.WikiAdmin,
		Label:       "知识库管理员",
		Description: "负责管理系统知识库内容的管理员",
	},
	{
		ID:          utils.BBSAdmin,
		Label:       "论坛管理员",
		Description: "负责管理论坛版块和帖子的管理员",
	},
	{
		ID:          utils.GalleryAdmin,
		Label:       "图库管理员",
		Description: "负责管理图片相册和图库的管理员",
	},
	{
		ID:          utils.ReviewAdmin,
		Label:       "评论管理员",
		Description: "负责管理用户评论和审核的管理员",
	},
}
