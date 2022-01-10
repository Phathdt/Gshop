package userrepo

import (
	"context"

	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (uint32, error)
	GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.User, error)
}

type createUserRepo struct {
	store CreateUserStorage
}

func NewCreateUserRepo(store CreateUserStorage) *createUserRepo {
	return &createUserRepo{store: store}
}

func (repo *createUserRepo) CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	if user, _ := repo.store.GetUserByCondition(ctx, map[string]interface{}{"username": data.Username}); user != nil {
		return nil, sdkcm.ErrCustom(nil, common.ErrExistedUser)
	}

	userId, err := repo.store.CreateUser(ctx, data)
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCreateUser)
	}

	user, err := repo.store.GetUserByCondition(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrRecordNotFound)
	}

	return user, nil
}
