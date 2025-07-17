package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func main() {

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("failed to load env: %v\n", err)
	}

	r := gin.Default()

	for _, service := range config.Services {
		routes, err := GetRoutesByService(service)
		if err != nil {
			fmt.Printf("failed to get routes from %s: %v\n", service, err)
			continue
		}
		proxyHandler := NewProxyHandler(service)
		for _, route := range routes {
			r.Handle(route.Method, route.Path, proxyHandler)
		}
	}

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
