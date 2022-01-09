package fibercart

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/carts/cartmodel"
	"gshop/module/carts/cartrepo"
	"gshop/module/carts/cartusecase"
	"gshop/module/products/productrepo"
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

		repo := cartrepo.NewCartRepo(sc.DB)
		productRepo := productrepo.NewProductRepo(sc.DB)
		uc := cartusecase.NewAddCartUseCase(repo, productRepo)

		if err := uc.AddToCart(c.Context(), user.ID, input.ProductId, input.Quantity); err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
