package main

import "github.com/axogc/backend/utils"

type Config struct {
	JWTKey string `envconfig:"JWT_KEY"`
	Port   string `envconfig:"PORT"`
	MySQL  utils.MySQLConfig
}
