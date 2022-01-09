package usrusecase

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/spf13/viper"
	"gshop/common"
	"gshop/module/users"
	"gshop/module/users/usrmodel"
	"gshop/sdk/sdkcm"
)

type userUseCase struct {
	Repo users.UserRepo
}

func (u userUseCase) CreateUser(ctx context.Context, input *usrmodel.UserCreate) (*usrmodel.User, error) {
	user, err := u.Repo.GetUserByUsername(ctx, input.Username)

	if user != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrExistedUser)
	}

	pwd := sha1.New()
	pwd.Write([]byte(input.Password))
	pwd.Write([]byte(viper.GetString("HASH_SALT")))
	input.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	if err = u.Repo.CreateUser(ctx, input); err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCreateUser)
	}

	return u.Repo.GetUserByUsername(ctx, input.Username)
}

func (u userUseCase) GetUserByUsername(ctx context.Context, username string) (*usrmodel.User, error) {
	return u.Repo.GetUserByUsername(ctx, username)
}

func NewUserUseCase(repo users.UserRepo) *userUseCase {
	return &userUseCase{Repo: repo}
}
