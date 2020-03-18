package v1

import (
	"net/http"
	"sixth-io/kaeru/db"
	"sixth-io/kaeru/hash"

	"github.com/labstack/echo/v4"
)

// Login struct implement AuthHandler type
type Login struct {
	db *db.Database
}

// LoginCredentials holds the login credentials
type LoginCredentials struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

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
	creds := new(LoginCredentials)
	if err := c.Bind(creds); err != nil {
		return err
	}
	user, err := l.db.FetchUserForEmail(creds.Email)
	if err != nil {
		return err
	}

	match, err := hash.Compare(creds.Password, user.PasswordHash)
	if err != nil {
		return err
	}
	if !match {
		return c.String(http.StatusForbidden, "Invalid credentials")
	}

	return c.String(http.StatusOK, user.ID)
}

// NewLogin returns an instance of the signup route
func NewLogin(db *db.Database) *Login {
	return &Login{
		db: db,
	}
}
