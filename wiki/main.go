package main

import (
	"log"

	"github.com/axogc/backend/utils"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	db, err := utils.InitMySQL(&config.MySQL, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v\n", err)
	}

	if err := GetRouter(&config, db).Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
