package v1

import (
	"net/http"
	"sixth-io/kaeru/db"
	"sixth-io/kaeru/hash"
	"sixth-io/kaeru/models"

	"github.com/labstack/echo/v4"
)

// Signup implementation struct
type Signup struct {
	db *db.Database
}

// SignupCredentials holds the signup credentials
type SignupCredentials struct {
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
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
	creds := new(SignupCredentials)
	if err := c.Bind(creds); err != nil {
		return err
	}

	hash, salt, err := hash.Password(creds.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:         creds.Name,
		Email:        creds.Email,
		PasswordHash: hash,
		Salt:         salt,
	}

	if err := l.db.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, creds)
}

// NewSignup returns an instance of the signup route
func NewSignup(db *db.Database) *Signup {
	return &Signup{
		db: db,
	}
}
