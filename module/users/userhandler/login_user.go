package userhandler

import (
	"context"

	"gshop/module/users/usermodel"
)

type LoginUserRepo interface {
	LoginUser(ctx context.Context, data *usermodel.UserLogin) (*usermodel.User, error)
}

type loginUserHdl struct {
	repo LoginUserRepo
}

func NewLoginUserHdl(repo LoginUserRepo) *loginUserHdl {
	return &loginUserHdl{repo: repo}
}

func (h *loginUserHdl) Response(ctx context.Context, data *usermodel.UserLogin) (*usermodel.User, error) {
	return h.repo.LoginUser(ctx, data)
}
