package userstorage

import (
	"context"

	"gshop/common"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

func (s *userSQLStorage) CreateUser(ctx context.Context, data *usermodel.UserCreate) (uint32, error) {
	data.Password = common.GetHash([]byte(data.Password))

	if err := s.db.Create(&data).Error; err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return data.ID, nil
}
