package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

func GenerateJWT(username string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   username,
		Issuer:    "MoneyWise",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
