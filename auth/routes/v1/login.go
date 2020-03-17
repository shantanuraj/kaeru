package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login struct implement AuthHandler type
type Login struct{}

// Method to serve on
func (l *Login) Method() string {
	return http.MethodPost
}

// URL to handle
func (l *Login) URL() string {
	return "/login"
}

// Handler for managing login
func (l *Login) Handler() echo.HandlerFunc {
	return l.login
}

func (l *Login) login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}
