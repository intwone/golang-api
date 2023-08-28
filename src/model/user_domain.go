package model

import (
	"encoding/json"

	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/util"
)

type userDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
}

func (ud *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(ud)

	if err != nil {
		logger.Error("error during to marshal json", err, util.CreateJourneyField("UserDomain"))
		return "", err
	}

	return string(json), nil
}
