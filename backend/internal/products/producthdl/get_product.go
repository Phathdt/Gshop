package producthdl

import (
	"context"

	"gshop/internal/products/productmodel"
)

type GetProductRepo interface {
	GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error)
}

type getProductHdl struct {
	repo GetProductRepo
}

func NewGetProductHdl(repo GetProductRepo) *getProductHdl {
	return &getProductHdl{repo: repo}
}

func (h *getProductHdl) Response(ctx context.Context, id uint32) (*productmodel.Product, error) {
	return h.repo.GetProduct(ctx, id)
}
