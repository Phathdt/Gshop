package carthdl

import (
	"context"

	"gshop/module/products/producthdl"
)

type AddToCartRepo interface {
	DeleteCartProduct(ctx context.Context, cartId, productId uint32) error
	AddToCart(ctx context.Context, cartId, productId, quantity, price uint32) error
	UpdateTotalCart(ctx context.Context, cartId uint32) error
}

type addToCartHdl struct {
	repo        AddToCartRepo
	readRepo    GetCartRepo
	productRepo producthdl.GetProductRepo
}

func NewAddToCartHdl(repo AddToCartRepo, readRepo GetCartRepo, productRepo producthdl.GetProductRepo) *addToCartHdl {
	return &addToCartHdl{repo: repo, readRepo: readRepo, productRepo: productRepo}
}

func (h *addToCartHdl) Response(ctx context.Context, userId, productId, quantity uint32) error {
	cart, err := h.readRepo.GetCart(ctx, userId)
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
