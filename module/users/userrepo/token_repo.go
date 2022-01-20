package userrepo

import (
	"context"

	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

type TokenStorage interface {
	CreateToken(ctx context.Context, token string, userId uint32) error
	GetToken(ctx context.Context, userId uint32, secretToken string) (string, error)
}

type tokenRepo struct {
	store TokenStorage
}

func (r *tokenRepo) GetToken(ctx context.Context, userId uint32, secretToken string) (string, error) {
	token, err := r.store.GetToken(ctx, userId, secretToken)

	if err != nil {
		return "", sdkcm.ErrCustom(err, common.ErrRedis)
	}

	if token == "" {
		return "", sdkcm.ErrCustom(err, common.ErrGetToken)
	}

	return token, nil
}

func (r *tokenRepo) CreateToken(ctx context.Context, user *usermodel.User) (string, error) {
	jwt, err := common.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	err = r.store.CreateToken(ctx, jwt, user.ID)
	if err != nil {
		return "", sdkcm.ErrCustom(err, common.ErrCreateToken)
	}

	return jwt, nil
}

func NewTokenRepo(store TokenStorage) *tokenRepo {
	return &tokenRepo{store: store}
}
