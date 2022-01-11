package productrepo

import (
	"context"

	"gshop/common"
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type ProductStorage interface {
	GetProductByCondition(ctx context.Context, cond map[string]interface{}) (*productmodel.Product, error)
	ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging, moreKeys ...string) ([]productmodel.Product, error)
}

type productRepo struct {
	store ProductStorage
}

func NewProductRepo(store ProductStorage) *productRepo {
	return &productRepo{store: store}
}

func (r *productRepo) GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error) {
	product, err := r.store.GetProductByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCannotGetProduct)
	}

	return product, nil
}

func (r *productRepo) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error) {
	products, err := r.store.ListProduct(ctx, filter, paging, "Category")
	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCannotListProduct)
	}

	return products, nil
}
