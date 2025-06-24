package main

import "github.com/axogc/backend/utils"

type Config struct {
	Port   string            `envconfig:"PORT"`
	JWTKey string            `envconfig:"JWT_KEY" default:"8080"`
	MySQL  utils.MySQLConfig `envconfig:"MYSQL"`
}
