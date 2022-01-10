package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/carts/cartrepo"
	"gshop/module/carts/cartusecase"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func MyCart(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := common.GetCurrentUser(c)

		repo := cartrepo.NewCartRepo(sc.DB)
		uc := cartusecase.NewCartUseCase(repo)

		cart, err := uc.MyCart(c.Context(), user.ID)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse(cart))
	}
}
