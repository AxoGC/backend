package utils

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CorsMidWare = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
})

var MimeToExt = map[string]string{
	"image/jpeg": ".jpeg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

var ErrNotSupportedImageType = errors.New("not supported image file type")

func GetExtByFileHeader(fh *multipart.FileHeader) (string, error) {

	file, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read file header: %w", err)
	}

	mime := http.DetectContentType(buf)
	ext, exists := MimeToExt[mime]
	if !exists {
		return "", ErrNotSupportedImageType
	}

	return ext, nil
}

func RandomCode(length uint, withLetters bool) string {
	var charset string
	if withLetters {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	} else {
		charset = "0123456789"
	}

	var sb strings.Builder
	for i := uint(0); i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

type PaginateConfig struct {
	KeyPrefix       string
	PageKey         string
	PageSizeKey     string
	DefaultPage     int
	DefaultPageSize int
}

func Paginate(c *gin.Context, cfg *PaginateConfig) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := 1
		pageSize := 10
		pageKey := "page"
		pageSizeKey := "page_size"
		if cfg != nil {
			if cfg.DefaultPage != 0 {
				page = cfg.DefaultPage
			}
			if cfg.DefaultPageSize != 0 {
				pageSize = cfg.DefaultPageSize
			}
			if cfg.PageKey != "" {
				pageKey = cfg.PageKey
			}
			if cfg.PageSizeKey != "" {
				pageSizeKey = cfg.PageSizeKey
			}
			if cfg.KeyPrefix != "" {
				pageKey = cfg.KeyPrefix + "_" + pageKey
				pageSizeKey = cfg.KeyPrefix + "_" + pageSizeKey
			}
		}
		if value, err := strconv.Atoi(c.Query(pageKey)); err != nil {
			page = value
		}
		if value, err := strconv.Atoi(c.Query(pageSizeKey)); err != nil {
			pageSize = value
		}
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

type TxResp[T any] struct {
	Code int
	Data *T
}

func (tr *TxResp[T]) Error() string {
	return fmt.Sprintf("code: %d, data: %v", tr.Code, tr.Data)
}

func TxRes[T any](code int, data *T) error {
	return &TxResp[T]{Code: code, Data: data}
}

type Handler[T any] func(cfg *T) (string, string, gin.HandlerFunc)

func RegisterHandlers[T any](r *gin.Engine, cfg *T, handlers ...Handler[T]) {
	for _, handler := range handlers {
		r.Handle(handler(cfg))
	}
}
