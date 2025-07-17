package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func GetRoutes(r *gin.Engine) gin.HandlerFunc {
	routes := lo.Map(r.Routes(), func(route gin.RouteInfo, _ int) Route {
		return Route{route.Method, route.Path}
	})
	return func(c *gin.Context) {
		c.JSON(200, routes)
	}
}
