package carthdl

import (
	"context"

	"gshop/internal/carts/cartmodel"
)

type GetCartRepo interface {
	GetCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error)
}

type getCartHdl struct {
	repo GetCartRepo
}

func NewGetCartHdl(repo GetCartRepo) *getCartHdl {
	return &getCartHdl{repo: repo}
}

func (h *getCartHdl) Response(ctx context.Context, userId uint32) (*cartmodel.Cart, error) {
	return h.repo.GetCart(ctx, userId)
}
