package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func ListServers(cfg *HandlerConfig) (string, string, []gin.HandlerFunc) {
	return "GET", "/servers", []gin.HandlerFunc{
		func(c *gin.Context) {

			var servers []struct {
			}

			if err := cfg.DB.Model(new(utils.Server)).Find(&servers).Error; err != nil {

			}
		},
	}
}
