package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func AuthorDomaintoAuthorSchema(request domain.Author) *schema.Author {
	return &schema.Author{
		Name: request.Name,
	}
}

func AuthorCreateRequestToAuthorDomain(request web.AuthorCreateRequest) *domain.Author {
	return &domain.Author{
		Name: request.Name,
	}
}

func AuthorUpdateRequestToAuthorDomain(request web.AuthorUpdateRequest) *domain.Author {
	return &domain.Author{
		Name: request.Name,
	}
}
