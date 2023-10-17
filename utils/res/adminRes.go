package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func AdminDomainToAdminLoginResponse(user *domain.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func AdminSchemaToAdminDomain(user *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func AdminDomaintoAdminResponse(user *domain.Admin) web.AdminReponse {
	return web.AdminReponse{
		Id:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ConvertAdminResponse(users []domain.Admin) []web.AdminReponse {
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
