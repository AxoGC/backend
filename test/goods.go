package main

import "github.com/axogc/backend/utils"

var TestGoods = []utils.Good{
	{
		ID:     1,
		PropID: utils.BlindBox,
		Label:  "新手盲盒",
		Count:  100,
		Price:  50,
	},
	{
		ID:     2,
		PropID: utils.BlindBox,
		Label:  "精美盲盒",
		Count:  50,
		Price:  100,
	},
	{
		ID:     3,
		PropID: utils.BlindBox,
		Label:  "传说盲盒",
		Count:  20,
		Price:  250,
	},
}
