package model

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/intwone/golang-api/src/configuration/rest_err"
)

var (
	JWT_SECRET = "JWT_SECRET"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"id":  ud.id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		message := fmt.Sprintf("error trying to generate jwt, error: %s", err.Error())
		return "", rest_err.NewInternalServerError(message)
	}

	return tokenString, nil
}
