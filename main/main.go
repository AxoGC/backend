package main

import (
	"log"

	"github.com/axogc/backend/utils"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to load config:%v\n", err)
	}

	db, err := utils.InitMySQL(&config.MySQL)
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v\n", err)
	}

	if err := GetRouter(&config, db).Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
