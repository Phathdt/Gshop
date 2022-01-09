package middleware

import (
	"fmt"
	"net/http"

	"gshop/sdk"
	"gshop/sdk/sdkcm"

	"github.com/gofiber/fiber/v2"
)

func Recover(sc *sdk.ServiceContext) fiber.Handler {
	logger := sc.Logger("service")
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(sdkcm.AppError); ok {
					appErr.RootCause = appErr.RootError()
					logger.Errorln(appErr.RootCause)

					if appErr.RootCause != nil {
						appErr.Log = appErr.RootCause.Error()
					}

					c.Status(appErr.StatusCode).JSON(&fiber.Map{
						"error": appErr,
					})

				} else {
					var appErr sdkcm.AppError

					if e, ok := err.(error); ok {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
						logger.Errorln(e.Error())

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

					} else {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
						logger.Errorln(fmt.Sprintf("%s", err))

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

					}
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
