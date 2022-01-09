package users

import (
	"context"
	"gshop/module/users/usrmodel"
)

type UserUseCase interface {
	GetUserByUsername(ctx context.Context, username string) (*usrmodel.User, error)
}
