package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Signup implementation struct
type Signup struct{}

// Method to serve on
func (l *Signup) Method() string {
	return http.MethodPost
}

// URL to server for signing up
func (l *Signup) URL() string {
	return "/signup"
}

// Handler for managing signup
func (l *Signup) Handler() echo.HandlerFunc {
	return l.signup
}

func (l *Signup) signup(c echo.Context) error {
	return c.String(http.StatusOK, "signup")
}
