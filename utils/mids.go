package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HandlerFunc[T any] func(c *gin.Context, u *User, r *T) (int, *Resp)

func WithCaptchaAuth(rdb *redis.Client, hf gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			Captcha     string `json:"captcha" binding:"required"`
			CaptchaCode string `json:"captchaCode" binding:"required"`
		}

		if err := c.ShouldBindBodyWithJSON(&request); err != nil {
			c.JSON(400, Resp{"用户请求有误", nil})
			c.Error(err)
			return
		}

		captchaCode, err := rdb.Get(context.Background(), "auth:captcha:"+request.Captcha).Result()

		if errors.Is(err, redis.Nil) {
			c.AbortWithStatusJSON(404, Resp{"验证码未申请", nil})
			return
		} else if err != nil {
			c.AbortWithStatusJSON(500, Resp{"获取验证码失败", nil})
			c.Error(err)
			return
		}

		if captchaCode != request.CaptchaCode {
			c.AbortWithStatusJSON(400, Resp{"验证码错误", nil})
			return
		}

		hf(c)
	}
}

func WithEmailAuth(rdb *redis.Client, hf gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			Email     string `json:"email" binding:"required,email"`
			EmailCode string `json:"emailCode" binding:"required,len=6"`
		}

		if err := c.ShouldBindBodyWithJSON(&request); err != nil {
			c.JSON(400, Resp{"用户请求有误", nil})
			c.Error(err)
			return
		}

		emailCode, err := rdb.Get(context.Background(), "auth:email:"+request.Email).Result()

		if errors.Is(err, redis.Nil) {
			c.AbortWithStatusJSON(404, Resp{"验证码未申请", nil})
			return
		} else if err != nil {
			c.AbortWithStatusJSON(500, Resp{"获取验证码失败", nil})
			c.Error(err)
			return
		}

		if emailCode != request.EmailCode {
			c.AbortWithStatusJSON(400, Resp{"验证码错误", nil})
			return
		}

		hf(c)
	}
}

func UploadImageMidWare(baseFolderPath, folderName, fileName string) func(c *gin.Context) (int, *Resp) {
	return func(c *gin.Context) (int, *Resp) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.Error(err)
			return 400, &Resp{"文件上传失败", nil}
		}

		ext, err := GetExtByFileHeader(fileHeader)
		if errors.Is(err, ErrNotSupportedImageType) {
			return 400, Res("文件类型有误", nil)
		} else if err != nil {
			c.Error(err)
			return 500, Res("读取文件失败", nil)
		}

		folderPath := filepath.Join(baseFolderPath, folderName)
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			c.JSON(500, Resp{"无法创建保存目录", nil})
			c.Error(err)
			return 500, &Resp{"无法创建保存目录", nil}
		}

		fp := filepath.Join(folderPath, fileName+ext)
		if err := c.SaveUploadedFile(fileHeader, fp); err != nil {
			c.Error(err)
			return 500, &Resp{"文件保存失败", nil}
		}

		return 200, &Resp{"文件上传成功", nil}
	}
}

func LogMidWare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		var userId *uint

		if idAny, exists := c.Get("userId"); exists {
			if id, ok := idAny.(uint); ok {
				userId = &id
			}
		}

		for _, err := range c.Errors.Errors() {
			log := Log{
				Path:    c.Request.URL.Path,
				Method:  c.Request.Method,
				Status:  c.Writer.Status(),
				Message: err,
				UserID:  userId,
			}
			if err := db.Create(&log).Error; err != nil {
				fmt.Printf("Failed to log: %+v, because %v\n", &log, err)
			} else {
				fmt.Printf("Log successfully: %+v\n", &log)
			}
		}
	}
}

func WithRolesAuth[T any](roles []RoleID, hf HandlerFunc[T]) HandlerFunc[T] {
	return func(c *gin.Context, u *User, r *T) (int, *Resp) {

		if !u.HasAnyRole(roles...) {
			return 403, Res("权限不足，访问被拒绝", nil)
		}

		return hf(c, u, r)
	}
}
