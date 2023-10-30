package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func AuthorSchematoAuthorDomain(author *schema.Author) *domain.Author {
	return &domain.Author{
		ID:   author.ID,
		Name: author.Name,
	}
}

func AuthorDomaintoAuthorResponse(author *domain.Author) web.AuthorReponse {
	return web.AuthorReponse{
		ID:   author.ID,
		Name: author.Name,
	}
}

func UpdateAuthorDomaintoAuthorResponse(id uint, author *domain.Author) web.AuthorReponse {
	return web.AuthorReponse{
		ID:   id,
		Name: author.Name,
	}
}

func ConvertAuthorResponse(authors []domain.Author) []web.AuthorReponse {
	var results []web.AuthorReponse
	for _, author := range authors {
		authorResponse := web.AuthorReponse{
			ID:   author.ID,
			Name: author.Name,
		}
		results = append(results, authorResponse)
	}
	return results
}
