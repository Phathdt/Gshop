package users

import (
	"context"

	"gshop/module/users/usrmodel"
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*usrmodel.User, error)
	CreateUser(ctx context.Context, input *usrmodel.UserCreate) error
}
