package carts

import (
	"context"

	"gshop/module/carts/cartmodel"
)

type CartUseCase interface {
	MyCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error)
	AddToCart(ctx context.Context, userId, productId, quantity uint32) error
	ClearMyCart(ctx context.Context, userId uint32) error
}

type CartRepo interface {
	MyCart(ctx context.Context, userId uint32, moreKeys ...string) (*cartmodel.Cart, error)
	DeleteCart(ctx context.Context, cartId uint32) error
	CreateCart(ctx context.Context, userId uint32) error
	DeleteCartProduct(ctx context.Context, cartId, productId uint32) error
	UpdateTotalCart(ctx context.Context, cartId uint32) error
	AddToCart(ctx context.Context, cartId uint32, productId, quantity, price uint32) error
}
