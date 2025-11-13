package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
	status := client.Ping(ctx)
	if err := status.Err(); err != nil {
		log.Printf("failed to connect to client:%s", err)
		return nil
	}
	return client
}
