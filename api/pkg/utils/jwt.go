package utils

import (
	"errors"
	"log"
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTSignedToken(token string) (bool, LoginClaims, error) {
	claims := LoginClaims{}

	if token == "" {
		return false, claims, nil
	}

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		key := []byte(os.Getenv("JWT_SECRET"))
		return key, nil
	})

	if err != nil {
		log.Println(token, err.Error())
		return false, LoginClaims{}, errors.New("error validating token")
	}

	return tkn.Valid, claims, nil
}
