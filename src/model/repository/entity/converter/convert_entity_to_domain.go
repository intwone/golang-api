package converter

import (
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository/entity"
)

func ConverterEntityToDomain(entity entity.UserEntity) *model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Name,
		entity.Password,
		entity.Age,
	)

	domain.SetId(entity.Id.Hex())

	return &domain
}
