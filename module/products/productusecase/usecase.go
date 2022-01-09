package productusecase

import (
	"context"

	"gshop/module/products"
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type productUseCase struct {
	Repo products.ProductRepo
}

func (p productUseCase) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error) {
	return p.Repo.ListProduct(ctx, filter, paging, "Category")
}

func NewProductUseCase(repo products.ProductRepo) *productUseCase {
	return &productUseCase{Repo: repo}
}
