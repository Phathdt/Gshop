package userrepo

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

type LoginUserStorage interface {
	GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.User, error)
}

type loginUserRepo struct {
	store LoginUserStorage
}

func NewLoginUserRepo(store LoginUserStorage) *loginUserRepo {
	return &loginUserRepo{store: store}
}

func (r loginUserRepo) LoginUser(ctx context.Context, data *usermodel.UserLogin) (*usermodel.User, error) {
	user, err := r.store.GetUserByCondition(ctx, map[string]interface{}{"username": data.Username})
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrRecordNotFound)
	}

	userPass := []byte(data.Password)
	dbPass := []byte(user.Password)

	if passErr := bcrypt.CompareHashAndPassword(dbPass, userPass); passErr != nil {
		return nil, sdkcm.ErrCustom(nil, common.ErrPasswordNotMatch)
	}

	return user, nil
}
