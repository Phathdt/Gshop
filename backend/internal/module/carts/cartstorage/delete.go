package cartstorage

import (
	"context"

	"gshop/internal/module/carts/cartmodel"
	"gshop/pkg/sdkcm"
)

func (s *cartSQLStorage) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	if err := s.db.Table(cartmodel.CartProduct{}.TableName()).Where("cart_id = ? and  product_id = ?", cartId, productId).
		Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (s *cartSQLStorage) DeleteCart(ctx context.Context, cartId uint32) error {
	if err := s.db.Table(cartmodel.Cart{}.TableName()).Where("id = ?", cartId).Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
