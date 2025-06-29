package main

import "github.com/axogc/backend/utils"

type Config struct {
	JWTKey          string            `envconfig:"JWT_KEY"`
	Port            string            `envconfig:"PORT"`
	ServerHost      string            `envconfig:"SERVER_HOST"`
	BedrockPassword string            `envconfig:"BEDROCK_PASSWORD"`
	MySQL           utils.MySQLConfig `envconfig:"MYSQL"`
	Redis           utils.RedisConfig `envconfig:"REDIS"`
}
