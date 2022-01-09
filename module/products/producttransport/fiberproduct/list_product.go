package fiberproduct

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/module/products/productmodel"
	"gshop/module/products/productrepo"
	"gshop/module/products/productusecase"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func ListProduct(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p productmodel.ListParam

		if err := c.QueryParser(&p); err != nil {
			panic(err)
		}

		p.FullFill()

		repo := productrepo.NewProductRepo(sc.DB)
		uc := productusecase.NewProductUseCase(repo)

		data, err := uc.ListProduct(c.Context(), p.ListFilter, &p.Paging)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.ResponseWithPaging(data, p.ListFilter, p.Paging))
	}
}
