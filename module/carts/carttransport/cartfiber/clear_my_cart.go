package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/carts/carthdl"
	"gshop/module/carts/cartrepo"
	"gshop/module/carts/cartstorage"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func ClearMyCart(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := common.GetCurrentUser(c)

		storage := cartstorage.NewCartSQLStorage(sc.DB)
		repo := cartrepo.NewClearCartRepo(storage)
		readRepo := cartrepo.NewGetCartRepo(storage)
		hdl := carthdl.NewClearCartHdl(repo, readRepo)

		err := hdl.Response(c.Context(), user.ID)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
