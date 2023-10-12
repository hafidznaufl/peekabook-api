package res

import (
	"peekabook/model/domain"
	"peekabook/model/web"
)

func AdminDomainToAdminLoginResponse(user *domain.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Name: user.Name,
		Email: user.Email,
	}
}

func ConvertResponseAdmin(users []domain.Admin) []web.AdminReponse {
	var results []web.AdminReponse
	for _, user := range users {
		userResponse := web.AdminReponse{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		results = append(results, userResponse)
	}
	return results
}
