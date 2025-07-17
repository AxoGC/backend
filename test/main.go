package main

import (
	"fmt"
	"log"

	"github.com/axogc/backend/utils"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	db, err := utils.InitMySQL(&config.MySQL)
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v\n", err)
	}

	datas := []any{
		TestRoles,
		TestProps,
		TestGames,
		TestGenders,
		TestGuildStatus,

		TestUsers,
		TestGuilds,
		TestDocGroups,
		TestForumGroups,

		TestGoods,
		TestServers,

		TestUserGuilds,
		//TestUserFollows,
		TestDonations,
		TestUserRoles,
		TestDocs,
		TestAlbums,
		TestFiles,
		TestReviews,
		TestUserProps,
		TestLogs,
		TestDeepSeekMessages,

		TestOnlines,
		TestImages,
		TestForums,

		TestPosts,
	}

	for _, value := range datas {
		if err := db.Save(value).Error; err != nil {
			log.Fatalf("Failed to save: %v\n", err)
		} else {
			fmt.Println("Successfully Saved.")
		}
	}

}
