package userstorage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gshop/pkg/sdkcm"
)

type tokenStore struct {
	client *redis.Client
}

func (s *tokenStore) DeleteTokenUsers(ctx context.Context, userId uint32) error {
	regex := fmt.Sprintf("users/%d/tokens/*", userId)

	keys, err := s.client.Keys(ctx, regex).Result()
	if err != nil {
		return sdkcm.ErrDB(err)
	}

	_, err = s.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			pipe.Del(ctx, key)
		}

		return nil
	})

	return nil
}

func (s *tokenStore) GetToken(ctx context.Context, userId uint32, secretToken string) (string, error) {
	key := fmt.Sprintf("users/%d/tokens/%s", userId, secretToken)

	value, err := s.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", sdkcm.ErrDB(err)
	}

	return value, nil
}

func (s *tokenStore) CreateToken(ctx context.Context, token string, userId uint32) error {
	tokens := strings.Split(token, ".")
	signature := tokens[len(tokens)-1]

	key := fmt.Sprintf("users/%d/tokens/%s", userId, signature)

	if err := s.client.Set(ctx, key, 1, time.Second*viper.GetDuration("TOKEN_TTL")).Err(); err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func NewTokenStore(client *redis.Client) *tokenStore {
	return &tokenStore{client: client}
}
