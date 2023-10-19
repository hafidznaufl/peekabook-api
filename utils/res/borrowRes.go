package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func BorrowSchematoBorrowDomain(borrow *schema.Borrow) *domain.Borrow {
	return &domain.Borrow{
		BookID: borrow.BookID,
		UserID: borrow.UserID,
		Date:   borrow.Date,
		Return: borrow.Return,
		Status: borrow.Status,
	}
}

func BorrowDomaintoBorrowResponse(borrow *domain.Borrow) web.BorrowResponse {
	return web.BorrowResponse{
		BookID: borrow.BookID,
		UserID: borrow.UserID,
		Date:   borrow.Date,
		Return: borrow.Return,
		Status: borrow.Status,
	}
}

func ConvertBorrowResponse(borrows []domain.Borrow) []web.BorrowResponse {
	var results []web.BorrowResponse
	for _, borrow := range borrows {
		borrowResponse := web.BorrowResponse{
			BookID: borrow.BookID,
			UserID: borrow.UserID,
			Date:   borrow.Date,
			Return: borrow.Return,
			Status: borrow.Status,
		}
		results = append(results, borrowResponse)
	}
	return results
}
