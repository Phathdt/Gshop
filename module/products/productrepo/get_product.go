package productrepo

import (
	"context"

	"gshop/common"
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type GetProductStorage interface {
	GetProductByCondition(ctx context.Context, cond map[string]interface{}) (*productmodel.Product, error)
}

type getProductRepo struct {
	store GetProductStorage
}

func NewGetProductRepo(store GetProductStorage) *getProductRepo {
	return &getProductRepo{store: store}
}

func (r *getProductRepo) GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error) {
	product, err := r.store.GetProductByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCannotGetProduct)
	}

	return product, nil
}
