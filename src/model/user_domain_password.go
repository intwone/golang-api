package model

import (
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/util"
)

func (ud *userDomain) EncryptPassword() {
	hashedPassword, err := util.HashPassword(ud.password)

	if err != nil {
		logger.Error("error during generate hash to password", err, util.CreateJourneyField("UserDomain"))
	}

	ud.password = hashedPassword
}
