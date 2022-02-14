package userhandler

import (
	"context"

	"gshop/internal/users/usermodel"
)

type GetUserRepo interface {
	GetUser(ctx context.Context, id uint32) (*usermodel.User, error)
}

type GetTokenRepo interface {
	GetToken(ctx context.Context, userId uint32, secret string) (string, error)
}

type getUserHdl struct {
	repo      GetUserRepo
	tokenRepo GetTokenRepo
}

func NewGetUserHdl(repo GetUserRepo, tokenRepo GetTokenRepo) *getUserHdl {
	return &getUserHdl{repo: repo, tokenRepo: tokenRepo}
}

func (h *getUserHdl) Response(ctx context.Context, userId uint32, secretToken string) (*usermodel.User, error) {
	_, err := h.tokenRepo.GetToken(ctx, userId, secretToken)

	if err != nil {
		return nil, err
	}

	return h.repo.GetUser(ctx, userId)
}
