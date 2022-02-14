package userhandler

import "context"

type DeleteTokenRepo interface {
	DeleteTokenUsers(ctx context.Context, userId uint32) error
}

type logoutUserHdl struct {
	tokenRepo DeleteTokenRepo
}

func NewLogoutUserHdl(tokenRepo DeleteTokenRepo) *logoutUserHdl {
	return &logoutUserHdl{tokenRepo: tokenRepo}
}

func (h *logoutUserHdl) Response(ctx context.Context, userId uint32) error {
	return h.tokenRepo.DeleteTokenUsers(ctx, userId)
}
