package main

import "github.com/axogc/backend/utils"

type Config struct {
	JWTKey   string            `envconfig:"JWT_KEY"`
	Port     string            `envconfig:"PORT" default:"8080"`
	FilePath string            `envconfig:"FILE_PATH"`
	MySQL    utils.MySQLConfig `envconfig:"MYSQL"`
}
