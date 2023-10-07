package res

import (
	"rent-app/model/domain"
	"rent-app/model/web"
)

func UserDomainToUserCreateResponse(user *domain.User) web.UserCreateResponse {
	return web.UserCreateResponse{
		Email: user.Email,
	}
}

func ConvertResponse(users []domain.User) []web.UserReponse {
	var results []web.UserReponse
	for _, user := range users {
		userResponse := web.UserReponse{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		results = append(results, userResponse)
	}
	return results
}
