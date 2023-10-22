package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func BorrowDomaintoBorrowSchema(request domain.Borrow) *schema.Borrow {
	return &schema.Borrow{
		BookID: request.BookID,
		UserID: request.UserID,
		Date:   request.Date,
		Return: request.Return,
		Status: request.Status,
	}
}

func BorrowCreateRequestToBorrowDomain(request web.BorrowCreateRequest) *domain.Borrow {
	return &domain.Borrow{
		BookID: request.BookID,
		UserID: request.UserID,
		Date:   request.Date,
		Return: request.Return,
		Status: request.Status,
	}
}

func BorrowUpdateRequestToBorrowDomain(request web.BorrowUpdateRequest) *domain.Borrow {
	return &domain.Borrow{
		BookID: request.BookID,
		UserID: request.UserID,
		Date:   request.Date,
		Return: request.Return,
		Status: request.Status,
	}
}
