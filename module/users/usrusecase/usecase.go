package usrusecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gshop/common"
	"gshop/module/users"
	"gshop/module/users/usrmodel"
	"gshop/sdk/sdkcm"
)

type userUseCase struct {
	Repo users.UserRepo
}

func (u userUseCase) GetUser(ctx context.Context, id uint32) (*usrmodel.User, error) {
	return u.Repo.GetUser(ctx, id)
}

func (u userUseCase) LoginUser(ctx context.Context, input *usrmodel.UserLogin) (*usrmodel.User, error) {
	user, err := u.Repo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrRecordNotFound)
	}

	userPass := []byte(input.Password)
	dbPass := []byte(user.Password)

	if passErr := bcrypt.CompareHashAndPassword(dbPass, userPass); passErr != nil {
		return nil, sdkcm.ErrCustom(nil, common.ErrPasswordNotMatch)
	}

	return user, nil
}

func (u userUseCase) CreateUser(ctx context.Context, input *usrmodel.UserCreate) (*usrmodel.User, error) {
	user, err := u.Repo.GetUserByUsername(ctx, input.Username)

	if user != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrExistedUser)
	}

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
