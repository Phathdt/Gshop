package cartrepo

import (
	"context"

	"gshop/common"
	"gshop/sdk/sdkcm"
)

type ClearCartStorage interface {
	DeleteCart(ctx context.Context, cartId uint32) error
}

type clearCartRepo struct {
	store ClearCartStorage
}

func NewClearCartRepo(store ClearCartStorage) *clearCartRepo {
	return &clearCartRepo{store: store}
}

func (r *clearCartRepo) ClearCart(ctx context.Context, cartId uint32) error {
	if err := r.store.DeleteCart(ctx, cartId); err != nil {
		return sdkcm.ErrCustom(err, common.ErrDeleteCart)
	}

	return nil
}
