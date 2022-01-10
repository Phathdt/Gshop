package cartstorage

import (
	"context"

	"gshop/module/carts/cartmodel"
	"gshop/sdk/sdkcm"
)

func (s *cartSQLStorage) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	if err := s.db.Table(cartmodel.CartProduct{}.TableName()).Where("cart_id = ? and  product_id = ?", cartId, productId).
		Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
