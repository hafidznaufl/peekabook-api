package helper

import (
	"rent-app/model/domain"
	"rent-app/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}