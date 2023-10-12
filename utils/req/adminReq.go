package req

import (
	"peekabook/model/domain"
	"peekabook/model/web"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Email: request.Email,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}