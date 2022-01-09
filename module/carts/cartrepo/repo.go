package cartrepo

import (
	"context"

	"gorm.io/gorm"
	"gshop/module/carts/cartmodel"
	"gshop/sdk/sdkcm"
)

type cartRepo struct {
	DB *gorm.DB
}

func (c cartRepo) DeleteCart(ctx context.Context, cartId uint32) error {
	//TODO implement me
	panic("implement me")
}

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

func (c cartRepo) UpdateTotalCart(ctx context.Context, cartId uint32) error {
	if err := c.DB.Exec(updateTotalCart, cartId).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (c cartRepo) CreateCart(ctx context.Context, userId uint32) error {
	newCart := cartmodel.Cart{
		Total:  0,
		UserId: userId,
	}

	if err := c.DB.Create(&newCart).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (c cartRepo) MyCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error) {
	var cart cartmodel.Cart

	db := c.DB.Table(cartmodel.Cart{}.TableName()).Preload("CartProduct")

	if err := db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}

		return nil, sdkcm.ErrDB(err)
	}

	return &cart, nil
}

func (c cartRepo) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	if err := c.DB.Table(cartmodel.CartProduct{}.TableName()).Where("cart_id = ? and  product_id = ?", cartId, productId).
		Delete(nil).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (c cartRepo) AddToCart(ctx context.Context, cartId uint32, productId, quantity, price uint32) error {
	item := cartmodel.CartProduct{
		Quantity:  quantity,
		Total:     price * quantity,
		CartId:    cartId,
		ProductId: productId,
	}

	if err := c.DB.Create(&item).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func NewCartRepo(DB *gorm.DB) *cartRepo {
	return &cartRepo{DB: DB}
}
