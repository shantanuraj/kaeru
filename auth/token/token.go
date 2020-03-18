package token

import (
	"sixth-io/kaeru/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// New returns a new JWT token
func New(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["admin"] = false
	claims["exp"] = oneMonthLater()

	return token.SignedString([]byte(user.PasswordHash))
}

func oneMonthLater() int64 {
	return time.Now().AddDate(0, 1, 0).Unix()
}
