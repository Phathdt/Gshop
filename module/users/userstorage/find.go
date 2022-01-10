package userstorage

import (
	"context"

	"gorm.io/gorm"
	"gshop/module/users/usermodel"
	"gshop/sdk/sdkcm"
)

func (s *userSQLStorage) GetUserByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.User, error) {
	var data usermodel.User

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}

		return nil, sdkcm.ErrDB(err)
	}

	return &data, nil
}
