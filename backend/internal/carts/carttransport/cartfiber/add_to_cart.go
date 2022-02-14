package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/carts/carthdl"
	"gshop/internal/carts/cartmodel"
	"gshop/internal/carts/cartrepo"
	"gshop/internal/carts/cartstorage"
	"gshop/internal/products/productrepo"
	"gshop/internal/products/productstorage"
	"gshop/pkg"
	"gshop/pkg/sdkcm"
)

func AddToCart(sc *pkg.ServiceContext) fiber.Handler {
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
