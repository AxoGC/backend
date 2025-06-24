package utils

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Host     string `envconfig:"HOST" default:"127.0.0.1"`
	Port     string `envconfig:"PORT" default:"3306"`
	DB       string `envconfig:"DB"`
}

func InitMySQL(config *MySQLConfig, opts ...gorm.Option) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB,
	)), opts...); err != nil {
		return nil, err
	}

	db.AutoMigrate(Tables...)

	return db, nil
}

type RedisConfig struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	Password string `envconfig:"PASSWORD"`
	DB       int    `envconfig:"DB"`
}

func InitRedis(config *RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
}
