package v1

import (
	"net/http"
	"sixth-io/kaeru/db"

	"github.com/labstack/echo/v4"
)

// Validate implementation struct
type Validate struct {
	db *db.Database
}

// Private specifies if the route should be validated for auth or not
func (l *Validate) Private() bool {
	return true
}

// Method to serve on
func (l *Validate) Method() string {
	return http.MethodGet
}

// URL to server for signing up
func (l *Validate) URL() string {
	return "/validate"
}

// Handler for managing Validate
func (l *Validate) Handler() echo.HandlerFunc {
	return l.validate
}

func (l *Validate) validate(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

// NewValidate returns an instance of the Validate route
func NewValidate(db *db.Database) *Validate {
	return &Validate{
		db: db,
	}
}
