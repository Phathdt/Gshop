package cartusecase

import (
	"context"

	"gshop/common"
	"gshop/module/carts"
	"gshop/module/carts/cartmodel"
	"gshop/module/products"
	"gshop/sdk/sdkcm"
)

type cartUseCase struct {
	Repo        carts.CartRepo
	productRepo products.ProductRepo
}

func NewAddCartUseCase(repo carts.CartRepo, productRepo products.ProductRepo) *cartUseCase {
	return &cartUseCase{
		Repo:        repo,
		productRepo: productRepo,
	}
}

func (c cartUseCase) MyCart(ctx context.Context, userId uint32) (*cartmodel.Cart, error) {
	cart, _ := c.Repo.MyCart(ctx, userId, "CartProduct.Product")

	if cart != nil {
		return cart, nil
	}

	if err := c.Repo.CreateCart(ctx, userId); err != nil {
		return nil, sdkcm.ErrCustom(err, common.ErrCreateCart)
	}

	return c.Repo.MyCart(ctx, userId)
}

func (c cartUseCase) AddToCart(ctx context.Context, userId, productId, quantity uint32) error {
	cart, _ := c.Repo.MyCart(ctx, userId)

	if cart == nil {
		_ = c.Repo.CreateCart(ctx, userId)

		newCart, err := c.Repo.MyCart(ctx, userId)
		if err != nil {
			return sdkcm.ErrCustom(err, common.ErrFindCart)
		}

		cart = newCart
	}

	product, err := c.productRepo.GetProduct(ctx, productId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrProductNotFound)
	}

	_ = c.Repo.DeleteCartProduct(ctx, cart.ID, productId)

	if err = c.Repo.AddToCart(ctx, cart.ID, productId, quantity, product.Price); err != nil {
		return sdkcm.ErrCustom(err, common.ErrAddToCart)
	}

	if err = c.Repo.UpdateTotalCart(ctx, cart.ID); err != nil {
		return sdkcm.ErrCustom(err, common.ErrAddToCart)
	}

	return nil
}

func (c cartUseCase) ClearMyCart(ctx context.Context, userId uint32) error {
	cart, err := c.Repo.MyCart(ctx, userId)
	if err != nil {
		return sdkcm.ErrCustom(err, common.ErrFindCart)
	}

	if err = c.Repo.DeleteCart(ctx, cart.ID); err != nil {
		return sdkcm.ErrCustom(err, common.ErrDeleteCart)
	}

	return nil
}

func NewCartUseCase(repo carts.CartRepo) *cartUseCase {
	return &cartUseCase{Repo: repo}
}
