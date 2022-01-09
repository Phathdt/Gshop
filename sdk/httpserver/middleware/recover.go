package middleware

import (
	"fmt"
	"net/http"

	"gshop/sdk"
	"gshop/sdk/sdkcm"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Recover(sc *sdk.ServiceContext) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(sdkcm.AppError); ok {
					appErr.RootCause = appErr.RootError()
					logrus.Errorln(appErr.RootCause)

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
						logrus.Errorln(e.Error())

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

					} else {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
						logrus.Errorln(fmt.Sprintf("%s", err))

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
