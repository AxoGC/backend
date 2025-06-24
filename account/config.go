package main

import (
	"github.com/axogc/backend/utils"
)

type Config struct {
	Port     string            `envconfig:"PORT" default:"8080"`
	JWTKey   string            `envconfig:"JWT_KEY"`
	FilePath string            `envconfig:"FILE_PATH"`
	MySQL    utils.MySQLConfig `envconfig:"MYSQL"`
	Redis    utils.RedisConfig `envconfig:"REDIS"`
}
