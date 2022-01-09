package products

import (
	"context"

	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type ProductUseCase interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error)
	GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error)
}

type ProductRepo interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging, moreKeys ...string) ([]productmodel.Product, error)
	GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error)
}
