package cartstorage

import (
	"context"

	"gshop/sdk/sdkcm"
)

const updateTotalCart = `-- name: UpdateTotalCart
UPDATE
	carts
SET
	total = t.total
FROM (
	SELECT
		cart_id,
		SUM(total) AS total
	FROM
		cart_products
	WHERE
		cart_id = $1
	GROUP BY
		cart_id) AS t
WHERE
	t.cart_id = carts.id
`

func (s *cartSQLStorage) UpdateTotalCart(ctx context.Context, cartId uint32) error {
	if err := s.db.Exec(updateTotalCart, cartId).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
