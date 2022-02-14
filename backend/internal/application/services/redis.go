package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gshop/internal/application/config"
)

func NewRedisService(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS.Host,
		DB:   cfg.REDIS.DB,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("rdb.Ping %w", err)
	}

	return rdb, nil
}
