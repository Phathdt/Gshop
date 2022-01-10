package userhandler

import (
	"context"

	"gshop/module/users/usermodel"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error)
}

type createUserHdl struct {
	repo CreateUserRepo
}

func NewCreateUserHdl(repo CreateUserRepo) *createUserHdl {
	return &createUserHdl{repo: repo}
}

func (h *createUserHdl) Response(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	return h.repo.CreateUser(ctx, data)
}
