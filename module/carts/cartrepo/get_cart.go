package cartrepo

import (
	"context"

	"gshop/common"
	"gshop/module/carts/cartmodel"
	"gshop/sdk/sdkcm"
)

type GetCartStorage interface {
	GetCartByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*cartmodel.Cart, error)
	CreateCart(ctx context.Context, userId uint32) (uint32, error)
}

type getCartRepo struct {
	store GetCartStorage
}

func NewGetCartRepo(store GetCartStorage) *getCartRepo {
	return &getCartRepo{store: store}
}

func (r *getCartRepo) GetCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error) {
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
