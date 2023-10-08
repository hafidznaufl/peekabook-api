package req

import (
	"rentabook/model/domain"
	"rentabook/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}