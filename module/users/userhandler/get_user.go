package userhandler

import (
	"context"

	"gshop/module/users/usermodel"
)

type GetUserRepo interface {
	GetUser(ctx context.Context, id uint32) (*usermodel.User, error)
}

type getUserHdl struct {
	repo GetUserRepo
}

func NewGetUserHdl(repo GetUserRepo) *getUserHdl {
	return &getUserHdl{repo: repo}
}

func (h *getUserHdl) Response(ctx context.Context, id uint32) (*usermodel.User, error) {
	return h.repo.GetUser(ctx, id)
}
