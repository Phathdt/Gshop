package productfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/internal/application/services"
	"gshop/internal/module/products/producthdl"
	"gshop/internal/module/products/productmodel"
	"gshop/internal/module/products/productrepo"
	"gshop/internal/module/products/productstorage"
	"gshop/pkg/sdkcm"
)

func ListProduct(sc *services.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p productmodel.ListParam

		if err := c.QueryParser(&p); err != nil {
			panic(err)
		}

		p.FullFill()

		storage := productstorage.NewProductSQLStorage(sc.DB)
		repo := productrepo.NewProductRepo(storage)
		hdl := producthdl.NewListProductHdl(repo)

		data, err := hdl.Response(c.Context(), p.ListFilter, &p.Paging)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.ResponseWithPaging(data, p.ListFilter, p.Paging))
	}
}
