package carthdl

import (
	"context"

	"gshop/module/carts/cartmodel"
)

type ClearCartRepo interface {
	ClearCart(ctx context.Context, cartId uint32) error
	GetCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error)
}

type clearCartHdl struct {
	repo ClearCartRepo
}

func NewClearCartHdl(repo ClearCartRepo) *clearCartHdl {
	return &clearCartHdl{repo: repo}
}

func (h *clearCartHdl) Response(ctx context.Context, userId uint32) error {
	cart, err := h.repo.GetCart(ctx, userId)
	if err != nil {
		return err
	}

	return h.repo.ClearCart(ctx, cart.ID)
}
