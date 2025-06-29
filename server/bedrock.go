package main

import (
	"github.com/axogc/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type BedrockCommand = utils.BedrockCommand

type BedrockResponse = utils.BedrockResponse

func GetBedrockCommands(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/bedrock/commands", func(c *gin.Context) {

		var req struct {
			Slug     string `json:"slug"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.Error(err)
			c.JSON(400, Res("请求格式有误", nil))
			return
		}

		if req.Password != cfg.Env.BedrockPassword {
			c.JSON(400, Res("密码不正确", nil))
			return
		}

		c.JSON(200, Res("", lo.Ternary(cfg.BCs[req.Slug] == nil, make([]BedrockCommand, 0), cfg.BCs[req.Slug])))
		cfg.BCs[req.Slug] = make([]BedrockCommand, 0)
		return
	}
}

func SetBedrockResponses(cfg *HandlerConfig) (string, string, gin.HandlerFunc) {
	return "POST", "/bedrock/responses", func(c *gin.Context) {

		var req struct {
			Slug     string `json:"slug"`
			Password string `json:"password"`
			ID       string `json:"id"`
			BedrockResponse
		}
		if err := c.BindJSON(&req); err != nil {
			c.Error(err)
			c.JSON(400, Res("请求格式有误", nil))
			return
		}

		if req.Password != cfg.Env.BedrockPassword {
			c.AbortWithStatus(400)
			c.JSON(400, Res("密码不正确", nil))
			return
		}

		if cfg.BRs[req.Slug] == nil {
			c.AbortWithStatus(400)
			c.JSON(400, Res("没有用户请求此内容", nil))
			return
		}

		cfg.BRs[req.Slug] <- req.BedrockResponse
		c.JSON(200, Res("", nil))
	}
}
