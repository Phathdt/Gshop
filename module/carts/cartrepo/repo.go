package cartrepo

import (
	"context"

	"gshop/common"
	"gshop/module/carts/cartmodel"
	"gshop/sdk/sdkcm"
)

type CartStorage interface {
	AddToCart(ctx context.Context, cartId uint32, productId, quantity, price uint32) error
	UpdateTotalCart(ctx context.Context, cartId uint32) error
	DeleteCartProduct(ctx context.Context, cartId, productId uint32) error
	DeleteCart(ctx context.Context, cartId uint32) error
	GetCartByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*cartmodel.Cart, error)
	CreateCart(ctx context.Context, userId uint32) (uint32, error)
}
type cartRepo struct {
	store CartStorage
}

func NewCartRepo(store CartStorage) *cartRepo {
	return &cartRepo{store: store}
}

func (r *cartRepo) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	err := r.store.DeleteCartProduct(ctx, cartId, productId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrDeleteCartProduct)
	}

	return nil
}

func (r *cartRepo) AddToCart(ctx context.Context, cartId, productId, quantity, price uint32) error {
	err := r.store.AddToCart(ctx, cartId, productId, quantity, price)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrAddToCart)
	}

	return nil
}

func (r *cartRepo) UpdateTotalCart(ctx context.Context, cartId uint32) error {
	err := r.store.UpdateTotalCart(ctx, cartId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrUpdateCart)
	}

	return nil
}

func (r *cartRepo) ClearCart(ctx context.Context, cartId uint32) error {
	if err := r.store.DeleteCart(ctx, cartId); err != nil {
		return sdkcm.ErrCustom(err, common.ErrDeleteCart)
	}

	return nil
}

func (r *cartRepo) GetCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error) {
	cart, _ := r.store.GetCartByCondition(ctx, map[string]interface{}{"user_id": userId}, "CartProduct.Product")

	if cart != nil {
		return cart, nil
	}

	cartId, err := r.store.CreateCart(ctx, userId)
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCreateCart)
	}

	newCart, err := r.store.GetCartByCondition(ctx, map[string]interface{}{"id": cartId}, "CartProduct.Product")

	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrFindCart)
	}

	return newCart, nil
}
