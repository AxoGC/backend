package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter(cfg *HandlerConfig) *gin.Engine {
	r := gin.Default()

	r.Use(utils.CorsMidWare)

	utils.RegisterHandlers(r, cfg,
		ListServers,
		GetServers,
		GetOnline,
	)

	r.GET("/routes", utils.GetRoutes(r))

	return r
}
