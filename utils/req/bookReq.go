package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func BookDomaintoBookSchema(request domain.Book) *schema.Book {
	return &schema.Book{
		Title:     request.Title,
		AuthorID:  request.AuthorID,
		Page:      request.Page,
		Years:     request.Years,
		Publisher: request.Publisher,
		Type:      request.Type,
		Quantity:  request.Quantity,
		Status:    request.Status,
	}
}

func BookCreateRequestToBookDomain(request web.BookCreateRequest) *domain.Book {
	return &domain.Book{
		Title:     request.Title,
		AuthorID:  request.AuthorID,
		Page:      request.Page,
		Years:     request.Years,
		Publisher: request.Publisher,
		Type:      request.Type,
		Quantity:  request.Quantity,
		Status:    request.Status,
	}
}

func BookUpdateRequestToBookDomain(request web.BookUpdateRequest) *domain.Book {
	return &domain.Book{
		Title:     request.Title,
		AuthorID:  request.AuthorID,
		Page:      request.Page,
		Years:     request.Years,
		Publisher: request.Publisher,
		Type:      request.Type,
		Quantity:  request.Quantity,
		Status:    request.Status,
	}
}
