package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, srvCfg *config.Server) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: srvCfg.RedisAddr,
	})
	status := client.Ping(ctx)
	if err := status.Err(); err != nil {
		log.Printf("failed to connect to client:%s", err)
		return nil
	}
	return client
}
