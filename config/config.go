package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB() {

	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

var RDB *redis.Client

// InitRedis 初始化并连接 Redis
func InitRedis() {

	// 配置 Redis 信息
	redisAddr := "localhost:6379"
	redisPassword := "123456"
	redisDB := 0

	// 初始化 Redis 客户端
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// 使用 Ping 命令检查 Redis 连接是否正常
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis successfully!")
	}
}
