package cartstorage

import (
	"context"

	cartmodel2 "gshop/internal/module/carts/cartmodel"
	"gshop/pkg/sdkcm"
)

func (s *cartSQLStorage) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	if err := s.db.Table(cartmodel2.CartProduct{}.TableName()).Where("cart_id = ? and  product_id = ?", cartId, productId).
		Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (s *cartSQLStorage) DeleteCart(ctx context.Context, cartId uint32) error {
	if err := s.db.Table(cartmodel2.Cart{}.TableName()).Where("id = ?", cartId).Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
