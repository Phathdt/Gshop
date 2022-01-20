package userstorage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gshop/sdk/sdkcm"
)

type tokenStore struct {
	client *redis.Client
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
