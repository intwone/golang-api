package converter

import (
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository/entity"
)

func ConverterDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
