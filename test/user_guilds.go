package main

import (
	"github.com/axogc/backend/utils"
)

var TestUserGuilds = []utils.UserGuild{
	{UserID: 1, UserGuildStatusID: utils.GuildAdmin, GuildID: 1},
	{UserID: 2, UserGuildStatusID: utils.GuildAdmin, GuildID: 1},
	{UserID: 3, UserGuildStatusID: utils.GuildAdmin, GuildID: 2},
	{UserID: 4, UserGuildStatusID: utils.GuildMember, GuildID: 2},
	{UserID: 5, UserGuildStatusID: utils.GuildAdmin, GuildID: 3},
	{UserID: 6, UserGuildStatusID: utils.GuildMember, GuildID: 3},
	{UserID: 7, UserGuildStatusID: utils.GuildMember, GuildID: 3},
	{UserID: 8, UserGuildStatusID: utils.GuildMember, GuildID: 3},
}
