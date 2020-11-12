package redisconn

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

var (
	CTX           = context.Background()
	client        *redis.Client
	redishost     = "127.0.0.1:6379"
	redispassword = "123456"
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: redispassword, // no password set
		DB:       0,             // use default DB
	})

	pong, err := client.Ping(CTX).Result()
	if err == nil {
		fmt.Println("redis client init  success", pong)
	} else {
		fmt.Println("redis client init error", err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}
