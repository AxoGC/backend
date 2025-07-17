package main

type Config struct {
	Port     string   `envconfig:"PORT" default:"8080"`
	Services []string `envconfig:"SERVICES"`
}
