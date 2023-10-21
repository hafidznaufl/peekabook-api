package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func BookSchematoBookDomain(book *schema.Book) *domain.Book {
	return &domain.Book{
		ID:        book.ID,
		Title:     book.Title,
		AuthorID:  book.AuthorID,
		Page:      book.Page,
		Publisher: book.Publisher,
		Type:      book.Type,
		Quantity:  book.Quantity,
		Status:    book.Status,
	}
}

func BookDomaintoBookResponse(book *domain.Book) web.BookResponse {
	return web.BookResponse{
		ID:         book.ID,
		Title:      book.Title,
		AuthorName: book.AuthorName,
		Page:       book.Page,
		Years:      book.Years,
		Publisher:  book.Publisher,
		Type:       book.Type,
		Quantity:   book.Quantity,
		Status:     book.Status,
	}
}

func CreateBookDomaintoBookResponse(book *domain.Book) web.CreateBookResponse {
	return web.CreateBookResponse{
		ID:        book.ID,
		Title:     book.Title,
		AuthorID:  book.AuthorID,
		Page:      book.Page,
		Years:     book.Years,
		Publisher: book.Publisher,
		Type:      book.Type,
		Quantity:  book.Quantity,
		Status:    book.Status,
	}
}

func UpdateBookDomaintoBookResponse(book *domain.Book) web.UpdateBookResponse {
	return web.UpdateBookResponse{
		ID:        book.ID,
		Title:     book.Title,
		AuthorID:  book.AuthorID,
		Page:      book.Page,
		Years:     book.Years,
		Publisher: book.Publisher,
		Type:      book.Type,
		Quantity:  book.Quantity,
		Status:    book.Status,
	}
}

func ConvertBookResponse(books []domain.Book) []web.BookResponse {
	var results []web.BookResponse
	for _, book := range books {
		bookResponse := web.BookResponse{
			ID:         book.ID,
			Title:      book.Title,
			AuthorName: book.AuthorName,
			Page:       book.Page,
			Years:      book.Years,
			Publisher:  book.Publisher,
			Type:       book.Type,
			Quantity:   book.Quantity,
			Status:     book.Status,
		}
		results = append(results, bookResponse)
	}
	return results
}
