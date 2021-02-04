package redisconn

import (
	"context"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"github.com/skydrive/logger"
)

var (
	CTX           = context.Background()
	client        *redis.Client
	redishost     = "127.0.0.1:6379"
	redispassword = "123456"
)

func Setup() {
	client = redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: redispassword, // no password set
		DB:       0,             // use default DB
	})

	pong, err := client.Ping(CTX).Result()
	if err == nil {
		logger.Info("redis client init  success", pong)
	} else {
		logger.Error("redis client init error", err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}
