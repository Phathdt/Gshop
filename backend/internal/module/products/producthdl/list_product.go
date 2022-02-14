package producthdl

import (
	"context"

	productmodel2 "gshop/internal/module/products/productmodel"
	"gshop/pkg/sdkcm"
)

type ListProductRepo interface {
	ListProduct(ctx context.Context, filter *productmodel2.ListFilter, paging *sdkcm.Paging) ([]productmodel2.Product, error)
}
type listProductHdl struct {
	repo ListProductRepo
}

func NewListProductHdl(repo ListProductRepo) *listProductHdl {
	return &listProductHdl{repo: repo}
}

func (h *listProductHdl) Response(ctx context.Context, filter *productmodel2.ListFilter, paging *sdkcm.Paging) ([]productmodel2.Product, error) {
	return h.repo.ListProduct(ctx, filter, paging)
}
