package producthdl

import (
	"context"

	"gshop/internal/products/productmodel"
	"gshop/pkg/sdkcm"
)

type ListProductRepo interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error)
}
type listProductHdl struct {
	repo ListProductRepo
}

func NewListProductHdl(repo ListProductRepo) *listProductHdl {
	return &listProductHdl{repo: repo}
}

func (h *listProductHdl) Response(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error) {
	return h.repo.ListProduct(ctx, filter, paging)
}
