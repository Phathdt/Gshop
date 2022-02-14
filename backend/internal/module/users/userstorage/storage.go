package userstorage

import (
	"context"

	"gorm.io/gorm"
	"gshop/common"
	usermodel2 "gshop/internal/module/users/usermodel"
	"gshop/pkg/sdkcm"
)

type userSQLStorage struct {
	db *gorm.DB
}

func NewUserSQLStorage(db *gorm.DB) *userSQLStorage {
	return &userSQLStorage{db: db}
}

func (s *userSQLStorage) CreateUser(ctx context.Context, data *usermodel2.UserCreate) (uint32, error) {
	data.Password = common.GetHash([]byte(data.Password))

	if err := s.db.Create(&data).Error; err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return data.ID, nil
}

func (s *userSQLStorage) GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel2.User, error) {
	var data usermodel2.User

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}

		return nil, sdkcm.ErrDB(err)
	}

	return &data, nil
}
