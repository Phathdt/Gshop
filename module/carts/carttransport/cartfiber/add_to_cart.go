package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/carts/carthdl"
	"gshop/module/carts/cartmodel"
	"gshop/module/carts/cartrepo"
	"gshop/module/carts/cartstorage"
	"gshop/module/products/productrepo"
	"gshop/module/products/productstorage"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func AddToCart(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input cartmodel.AddToCartDTO
		if err := c.BodyParser(&input); err != nil {
			panic(err)
		}

		user := common.GetCurrentUser(c)

		storage := cartstorage.NewCartSQLStorage(sc.DB)
		readRepo := cartrepo.NewGetCartRepo(storage)
		repo := cartrepo.NewAddToCartRepo(storage)

		productStorage := productstorage.NewProductSQLStorage(sc.DB)
		productRepo := productrepo.NewGetProductRepo(productStorage)

		hdl := carthdl.NewAddToCartHdl(repo, readRepo, productRepo)

		if err := hdl.Response(c.Context(), user.ID, input.ProductId, input.Quantity); err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
