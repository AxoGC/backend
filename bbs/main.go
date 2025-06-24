package main

import (
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
		log.Fatalf("Failed to init database: %v\n", err)
	}

	r := GetRouter(&config, db)

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}
