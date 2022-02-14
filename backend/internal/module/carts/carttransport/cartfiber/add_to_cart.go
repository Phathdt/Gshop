package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/application/services"
	"gshop/internal/module/carts/carthdl"
	"gshop/internal/module/carts/cartmodel"
	"gshop/internal/module/carts/cartrepo"
	"gshop/internal/module/carts/cartstorage"
	"gshop/internal/module/products/productrepo"
	"gshop/internal/module/products/productstorage"
	"gshop/pkg/sdkcm"
)

func AddToCart(sc *services.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input cartmodel.AddToCartDTO
		if err := c.BodyParser(&input); err != nil {
			panic(err)
		}

		user := common.GetCurrentUser(c)

		storage := cartstorage.NewCartSQLStorage(sc.DB)
		repo := cartrepo.NewCartRepo(storage)

		productStorage := productstorage.NewProductSQLStorage(sc.DB)
		productRepo := productrepo.NewProductRepo(productStorage)

		hdl := carthdl.NewAddToCartHdl(repo, productRepo)

		if err := hdl.Response(c.Context(), user.ID, input.ProductId, input.Quantity); err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
