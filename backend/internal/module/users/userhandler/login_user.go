package userhandler

import (
	"context"

	usermodel2 "gshop/internal/module/users/usermodel"
)

type LoginUserRepo interface {
	LoginUser(ctx context.Context, data *usermodel2.UserLogin) (*usermodel2.User, error)
}

type loginUserHdl struct {
	repo      LoginUserRepo
	tokenRepo CreateTokenRepo
}

func NewLoginUserHdl(repo LoginUserRepo, tokenRepo CreateTokenRepo) *loginUserHdl {
	return &loginUserHdl{repo: repo, tokenRepo: tokenRepo}
}

func (h *loginUserHdl) Response(ctx context.Context, data *usermodel2.UserLogin) (string, error) {
	user, err := h.repo.LoginUser(ctx, data)

	if err != nil {
		return "", err
	}

	return h.tokenRepo.CreateToken(ctx, user)
}
