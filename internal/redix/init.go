package redix

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("REDIS_HOST"),
		DB:   viper.GetInt("REDIS_DB"),
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("rdb.Ping %w", err)
	}

	return rdb, nil
}
