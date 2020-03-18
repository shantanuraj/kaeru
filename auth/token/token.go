package token

import (
	"sixth-io/kaeru/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims are custom claims extending default ones.
type Claims struct {
	jwt.StandardClaims
}

// New returns a new JWT token
func New(user models.User) (string, error) {
	claims := &Claims{
		jwt.StandardClaims{
			Id:        user.ID,
			ExpiresAt: getExpiry(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(user.PasswordHash))
}

func getExpiry() int64 {
	return oneMonthLater()
}

func oneMonthLater() int64 {
	return time.Now().AddDate(0, 1, 0).Unix()
}
