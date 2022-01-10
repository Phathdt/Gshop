package carthdl

import "context"

type ClearCartRepo interface {
	ClearCart(ctx context.Context, cartId uint32) error
}

type clearCartHdl struct {
	repo     ClearCartRepo
	readRepo GetCartRepo
}

func NewClearCartHdl(repo ClearCartRepo, readRepo GetCartRepo) *clearCartHdl {
	return &clearCartHdl{repo: repo, readRepo: readRepo}
}

func (h *clearCartHdl) Response(ctx context.Context, userId uint32) error {
	cart, err := h.readRepo.GetCart(ctx, userId)
	if err != nil {
		return err
	}

	return h.repo.ClearCart(ctx, cart.ID)
}
