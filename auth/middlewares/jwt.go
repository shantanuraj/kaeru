package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const authorization = "Authorization"

// VerifyToken is the middleware function.
func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get(echo.HeaderAuthorization)

		if len(auth) == 0 {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Missing Authorization header",
			})
		}

		return next(c)
	}
}
