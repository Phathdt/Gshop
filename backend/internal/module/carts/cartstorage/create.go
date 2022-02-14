package cartstorage

import (
	"context"

	cartmodel2 "gshop/internal/module/carts/cartmodel"
	"gshop/pkg/sdkcm"
)

func (s cartSQLStorage) CreateCart(ctx context.Context, userId uint32) (uint32, error) {
	newCart := cartmodel2.Cart{UserId: userId}

	if err := s.db.Create(&newCart).Error; err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return newCart.ID, nil
}

func (s *cartSQLStorage) AddToCart(ctx context.Context, cartId uint32, productId, quantity, price uint32) error {
	item := cartmodel2.CartProduct{
		Quantity:  quantity,
		Total:     price * quantity,
		CartId:    cartId,
		ProductId: productId,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
