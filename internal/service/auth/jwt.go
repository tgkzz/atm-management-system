package auth

import (
	"atm/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

func createJwtToken(user models.User) (string, error) {
	claims := models.JWTClaims{
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Id,
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte("my-super-secret-secret-key"))
	if err != nil {
		return "", err
	}

	return token, err
}
