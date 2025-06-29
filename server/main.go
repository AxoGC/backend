package main

import (
	"log"

	"github.com/axogc/backend/utils"
	p "github.com/bestcb2333/gin-gorm-preloader/v2"
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

	if err := GetRouter(&HandlerConfig{
		Config: &p.Config{
			DB:     db,
			JWTKey: config.JWTKey,
		},
		Env: &config,
		RDB: utils.InitRedis(&config.Redis),
		BCs: make(map[string][]BedrockCommand),
		BRs: make(map[string]map[string]chan BedrockResponse),
	}).Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
