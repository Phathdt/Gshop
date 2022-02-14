package userrepo

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

type UserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (uint32, error)
	GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.User, error)
}

type userRepo struct {
	store UserStorage
}

func (r *userRepo) GetUser(ctx context.Context, id uint32) (*usermodel.User, error) {
	user, err := r.store.GetUserByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrFindUser)
	}

	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	if user, _ := r.store.GetUserByCondition(ctx, map[string]interface{}{"username": data.Username}); user != nil {
		return nil, sdkcm.ErrCustom(nil, common.ErrExistedUser)
	}

	userId, err := r.store.CreateUser(ctx, data)
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCreateUser)
	}

	user, err := r.store.GetUserByCondition(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrRecordNotFound)
	}

	return user, nil
}

func (r *userRepo) LoginUser(ctx context.Context, data *usermodel.UserLogin) (*usermodel.User, error) {
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

func NewUserRepo(store UserStorage) *userRepo {
	return &userRepo{store: store}
}
