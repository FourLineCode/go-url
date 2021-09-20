package utils

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type LoginClaims struct {
	jwt.StandardClaims

	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewJWTSignedToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTSignedToken(token string) (bool, error) {
	claims := LoginClaims{}

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil || !tkn.Valid {
		return false, errors.New("Invalid JWT token")
	}

	return true, nil
}
