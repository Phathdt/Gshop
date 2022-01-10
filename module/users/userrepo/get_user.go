package userrepo

import (
	"context"

	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

type GetUserStorage interface {
	GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.User, error)
}

type getUserRepo struct {
	store GetUserStorage
}

func NewGetUserRepo(store GetUserStorage) *getUserRepo {
	return &getUserRepo{store: store}
}

func (repo *getUserRepo) GetUser(ctx context.Context, id uint32) (*usermodel.User, error) {
	user, err := repo.store.GetUserByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrFindUser)
	}

	return user, nil
}
