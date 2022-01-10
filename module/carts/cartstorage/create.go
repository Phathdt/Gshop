package cartstorage

import (
	"context"

	"gshop/module/carts/cartmodel"
	"gshop/sdk/sdkcm"
)

func (s cartSQLStorage) CreateCart(ctx context.Context, userId uint32) (uint32, error) {
	newCart := cartmodel.Cart{UserId: userId}

	if err := s.db.Create(&newCart).Error; err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return newCart.ID, nil
}
