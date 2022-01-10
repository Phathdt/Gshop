package cartrepo

import (
	"context"

	"gshop/common"
	"gshop/sdk/sdkcm"
)

type AddToCartStorage interface {
	AddToCart(ctx context.Context, cartId uint32, productId, quantity, price uint32) error
	UpdateTotalCart(ctx context.Context, cartId uint32) error
	DeleteCartProduct(ctx context.Context, cartId, productId uint32) error
}

type addToCartRepo struct {
	store AddToCartStorage
}

func NewAddToCartRepo(store AddToCartStorage) *addToCartRepo {
	return &addToCartRepo{store: store}
}

func (r *addToCartRepo) DeleteCartProduct(ctx context.Context, cartId, productId uint32) error {
	err := r.store.DeleteCartProduct(ctx, cartId, productId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrDeleteCartProduct)
	}

	return nil
}

func (r *addToCartRepo) AddToCart(ctx context.Context, cartId, productId, quantity, price uint32) error {
	err := r.store.AddToCart(ctx, cartId, productId, quantity, price)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrAddToCart)
	}

	return nil
}

func (r *addToCartRepo) UpdateTotalCart(ctx context.Context, cartId uint32) error {
	err := r.store.UpdateTotalCart(ctx, cartId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrUpdateCart)
	}

	return nil
}
