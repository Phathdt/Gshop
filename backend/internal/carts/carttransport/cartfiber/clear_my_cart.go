package cartfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/carts/carthdl"
	"gshop/internal/carts/cartrepo"
	"gshop/internal/carts/cartstorage"
	"gshop/pkg"
	"gshop/pkg/sdkcm"
)

func ClearMyCart(sc *pkg.ServiceContext) fiber.Handler {
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
