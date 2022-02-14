package userhandler

import (
	"context"

	"gshop/internal/module/users/usermodel"
)

type LoginUserRepo interface {
	LoginUser(ctx context.Context, data *usermodel.UserLogin) (*usermodel.User, error)
}

type loginUserHdl struct {
	repo      LoginUserRepo
	tokenRepo CreateTokenRepo
}

func NewLoginUserHdl(repo LoginUserRepo, tokenRepo CreateTokenRepo) *loginUserHdl {
	return &loginUserHdl{repo: repo, tokenRepo: tokenRepo}
}

func (h *loginUserHdl) Response(ctx context.Context, data *usermodel.UserLogin) (string, error) {
	user, err := h.repo.LoginUser(ctx, data)

	if err != nil {
		return "", err
	}

	return h.tokenRepo.CreateToken(ctx, user)
}
