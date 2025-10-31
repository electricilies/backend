package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Cfg.RedisAddr,
	})
	status := client.Ping(context.Background())
	if err := status.Err(); err != nil {
		log.Printf("failed to connect to client:%s", err)
		return nil
	}
	return client
}
