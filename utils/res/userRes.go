package res

import (
	"peekabook/model/domain"
	"peekabook/model/web"
)

func UserDomainToUserLoginResponse(user *domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Name: user.Name,
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
