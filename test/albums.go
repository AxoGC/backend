package main

import (
	"github.com/axogc/backend/utils"
)

const (
	GalleryAlbumID = iota
	NewYear2023AlbumID
)

var TestAlbums = []utils.Album{
	{
		ID:          1,
		UserID:      1,
		GuildID:     nil,
		Slug:        "carousels",
		Label:       "主页轮播图",
		Profile:     "主页轮播图",
		Pinned:      true,
		Hide:        false,
		Protected:   true,
		ImageCount:  12,
		ReviewCount: 18,
	},
	{
		ID:          2,
		UserID:      1,
		GuildID:     nil,
		Slug:        "2023-new-year",
		Label:       "2023新年",
		Profile:     "2023新年的时候截的图",
		Pinned:      false,
		Hide:        false,
		Protected:   true,
		ImageCount:  25,
		ReviewCount: 9,
	},
}
