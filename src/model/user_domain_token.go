package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func VerifyToken(value string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	token, err := jwt.Parse(RemoveBearerPrefix(value), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		id: claims["id"].(string),
	}, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET)

	token, err := jwt.Parse(RemoveBearerPrefix(c.Request.Header.Get("Authorization")), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
