package usrrepo

import (
	"context"

	"gorm.io/gorm"
	"gshop/module/users/usrmodel"
	"gshop/sdk/sdkcm"
)

type userRepo struct {
	DB *gorm.DB
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
