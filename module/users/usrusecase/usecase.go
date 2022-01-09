package usrusecase

import (
	"context"

	"gshop/module/users"
	"gshop/module/users/usrmodel"
)

type userUseCase struct {
	Repo users.UserRepo
}

func (u userUseCase) GetUserByUsername(ctx context.Context, username string) (*usrmodel.User, error) {
	return u.Repo.GetUserByUsername(ctx, username)
}

func NewUserUseCase(repo users.UserRepo) *userUseCase {
	return &userUseCase{Repo: repo}
}
