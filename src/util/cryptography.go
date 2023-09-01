package util

import (
	"github.com/intwone/golang-api/src/configuration/logger"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error("error during generate hash to password", err, CreateJourneyField("HashPassword"))
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		logger.Error("error during compare password", err, CreateJourneyField("ComparePassword"))
		return false
	}

	return true
}
