package products

import (
	"context"

	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type ProductUseCase interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error)
}

type ProductRepo interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error)
}
