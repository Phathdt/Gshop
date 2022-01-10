package productrepo

import (
	"context"

	"gshop/common"
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type ListProductStorage interface {
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging, moreKeys ...string) ([]productmodel.Product, error)
}

type listProductRepo struct {
	store ListProductStorage
}

func NewListProductRepo(store ListProductStorage) *listProductRepo {
	return &listProductRepo{store: store}
}

func (r *listProductRepo) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error) {
	products, err := r.store.ListProduct(ctx, filter, paging, "Category")
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCannotListProduct)
	}

	return products, nil
}
