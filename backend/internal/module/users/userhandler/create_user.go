package userhandler

import (
	"context"

	"gshop/internal/module/users/usermodel"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error)
}

type CreateTokenRepo interface {
	CreateToken(ctx context.Context, user *usermodel.User) (string, error)
}

type createUserHdl struct {
	repo      CreateUserRepo
	tokenRepo CreateTokenRepo
}

func NewCreateUserHdl(repo CreateUserRepo, tokenRepo CreateTokenRepo) *createUserHdl {
	return &createUserHdl{repo: repo, tokenRepo: tokenRepo}
}

func (h *createUserHdl) Response(ctx context.Context, data *usermodel.UserCreate) (string, error) {
	user, err := h.repo.CreateUser(ctx, data)

	if err != nil {
		return "", err
	}

	return h.tokenRepo.CreateToken(ctx, user)
}
