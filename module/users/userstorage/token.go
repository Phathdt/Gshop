package userstorage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
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

func (s *tokenStore) CreateToken(ctx context.Context, secret, token string, userId uint32) error {
	tokens := strings.Split(token, ".")[0:2]
	newToken := strings.Join(tokens, ".")

	key := fmt.Sprintf("users/%d/tokens/%s", userId, secret)

	if err := s.client.Set(ctx, key, newToken, time.Minute*120).Err(); err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func NewTokenStore(client *redis.Client) *tokenStore {
	return &tokenStore{client: client}
}
