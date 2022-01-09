package usrrepo

import (
	"context"

	"gshop/common"
	"gshop/module/users/usrmodel"
	"gshop/sdk/sdkcm"

	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func (u userRepo) CreateUser(ctx context.Context, input *usrmodel.UserCreate) error {
	input.Password = common.GetHash([]byte(input.Password))

	if err := u.DB.Create(&input).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (u userRepo) GetUserByUsername(ctx context.Context, username string) (*usrmodel.User, error) {
	var user usrmodel.User

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}
		return nil, sdkcm.ErrDB(err)
	}

	return &user, nil
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{DB: DB}
}
