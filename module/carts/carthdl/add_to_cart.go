package carthdl

import (
	"context"

	"gshop/module/carts/cartmodel"
	"gshop/module/products/producthdl"
)

type AddToCartRepo interface {
	DeleteCartProduct(ctx context.Context, cartId, productId uint32) error
	AddToCart(ctx context.Context, cartId, productId, quantity, price uint32) error
	UpdateTotalCart(ctx context.Context, cartId uint32) error
	GetCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error)
}

type addToCartHdl struct {
	repo        AddToCartRepo
	productRepo producthdl.GetProductRepo
}

var ABC = 1

func NewAddToCartHdl(repo AddToCartRepo, productRepo producthdl.GetProductRepo) *addToCartHdl {
	ABC += 1

	return &addToCartHdl{repo: repo, productRepo: productRepo}
}

func (h *addToCartHdl) Response(ctx context.Context, userId, productId, quantity uint32) error {
	cart, err := h.repo.GetCart(ctx, userId)
	if err != nil {
		return err
	}

	product, err := h.productRepo.GetProduct(ctx, productId)
	if err != nil {
		return err
	}

	_ = h.repo.DeleteCartProduct(ctx, cart.ID, product.ID)

	if err = h.repo.AddToCart(ctx, cart.ID, product.ID, quantity, product.Price); err != nil {
		return err
	}

	if err = h.repo.UpdateTotalCart(ctx, cart.ID); err != nil {
		return err
	}

	return nil
}
