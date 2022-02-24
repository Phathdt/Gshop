package productfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/module/products/producthdl"
	"gshop/module/products/productmodel"
	"gshop/module/products/productrepo"
	"gshop/module/products/productstorage"
	"gshop/pkg/sdkcm"
	"gshop/svcctx"
)

func ListProduct(sc *svcctx.ServiceContext) fiber.Handler {
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
