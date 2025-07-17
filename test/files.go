package main

import (
	"time"

	"github.com/axogc/backend/utils"
)

var TestFiles = []utils.File{
	{
		ID:          1,
		CreatedAt:   time.Now().AddDate(0, -1, 0),
		Name:        "redstone_circuit_tutorial.zip",
		UserID:      1,
		Description: "红石电路教程相关的存档文件和原理图",
	},
	{
		ID:          2,
		CreatedAt:   time.Now().AddDate(0, 0, -20),
		Name:        "castle_blueprint.schematic",
		UserID:      1,
		Description: "中世纪城堡的建筑蓝图文件",
	},
	{
		ID:          3,
		CreatedAt:   time.Now().AddDate(0, 0, -10),
		Name:        "terraria_world_backup.wld",
		UserID:      1,
		Description: "泰拉瑞亚世界存档备份，包含完整boss战场地",
	},
}
