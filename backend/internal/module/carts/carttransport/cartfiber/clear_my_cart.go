package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/internal/application/services"
	"gshop/internal/module/carts/carthdl"
	"gshop/internal/module/carts/cartrepo"
	"gshop/internal/module/carts/cartstorage"
	"gshop/pkg/common"
	"gshop/pkg/sdkcm"
)

func ClearMyCart(sc *services.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := common.GetCurrentUser(c)

		storage := cartstorage.NewCartSQLStorage(sc.DB)
		repo := cartrepo.NewCartRepo(storage)
		hdl := carthdl.NewClearCartHdl(repo)

		err := hdl.Response(c.Context(), user.ID)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
