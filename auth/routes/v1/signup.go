package v1

import (
	"net/http"
	"sixth-io/kaeru/db"
	"sixth-io/kaeru/models"

	"github.com/labstack/echo/v4"
)

// Signup implementation struct
type Signup struct {
	db *db.Database
}

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
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := l.db.CreateUser(*u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}

// NewSignup returns an instance of the signup route
func NewSignup(db *db.Database) *Signup {
	return &Signup{
		db: db,
	}
}
