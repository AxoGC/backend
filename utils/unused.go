package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SendRespGet(c *gin.Context, result *gorm.DB, data any) {
	if err := result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, Resp{"找不到对应的数据", nil})
		return
	} else if err != nil {
		c.JSON(500, Resp{"查找失败", nil})
		c.Error(err)
		return
	}
	c.JSON(200, Resp{"查找成功", data})
}

func SendRespDelete(c *gin.Context, result *gorm.DB) {
	if result.RowsAffected == 0 {
		c.JSON(404, Resp{"找不到删除的数据", nil})
		return
	} else if result.Error != nil {
		c.JSON(500, Resp{"删除失败", nil})
		c.Error(result.Error)
		return
	}
	c.JSON(200, Resp{fmt.Sprintf("成功删除 %d 条数据", result.RowsAffected), nil})
}

func HandleQueryError(c *gin.Context, err error, msgNotFound, msgError string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, Resp{msgNotFound, nil})
	} else if err != nil {
		c.JSON(500, Resp{msgError, nil})
		c.Error(err)
	}
}

func Get[T any](url string) (*T, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
