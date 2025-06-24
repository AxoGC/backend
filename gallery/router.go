package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(config *Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	cfg := &p.Config{}

	RegFoldersHandler(r, cfg)
	RegFilesHandler(r, cfg)

	return r
}
