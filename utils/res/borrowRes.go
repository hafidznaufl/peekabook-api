package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func BorrowSchematoBorrowDomain(borrow *schema.Borrow) *domain.Borrow {
	return &domain.Borrow{
		ID:     borrow.ID,
		BookID: borrow.BookID,
		UserID: borrow.UserID,
		Date:   borrow.Date,
		Return: borrow.Return,
		Status: borrow.Status,
	}
}

func BorrowDomaintoBorrowResponse(borrow *domain.Borrow) web.BorrowResponse {
	return web.BorrowResponse{
		ID:        borrow.ID,
		BookTitle: borrow.BookTitle,
		UserName:  borrow.UserName,
		Date:      borrow.Date,
		Return:    borrow.Return,
		Status:    borrow.Status,
	}
}

func UpdateBorrowDomaintoBorrowResponse(id uint, borrow *domain.Borrow) web.BorrowResponse {
	return web.BorrowResponse{
		ID:        id,
		BookTitle: borrow.BookTitle,
		UserName:  borrow.UserName,
		Date:      borrow.Date,
		Return:    borrow.Return,
		Status:    borrow.Status,
	}
}

func CreateBorrowDomaintoBorrowResponse(borrow *domain.Borrow) web.CreateBorrowResponse {
	return web.CreateBorrowResponse{
		ID:     borrow.ID,
		BookID: borrow.BookID,
		UserID: borrow.UserID,
		Date:   borrow.Date,
		Return: borrow.Return,
		Status: borrow.Status,
	}
}

func ReturnBorrowDomaintoBorrowResponse(borrow *domain.Borrow) web.BorrowResponse {
	return web.BorrowResponse{
		ID:        borrow.ID,
		BookID:    borrow.BookID,
		BookTitle: borrow.BookTitle,
		UserName:  borrow.UserName,
		Date:      borrow.Date,
		Return:    borrow.Return,
		Status:    borrow.Status,
	}
}

func ConvertBorrowResponse(borrows []domain.Borrow) []web.BorrowResponse {
	var results []web.BorrowResponse
	for _, borrow := range borrows {
		borrowResponse := web.BorrowResponse{
			ID:        borrow.ID,
			BookTitle: borrow.BookTitle,
			UserName:  borrow.UserName,
			Date:      borrow.Date,
			Return:    borrow.Return,
			Status:    borrow.Status,
		}
		results = append(results, borrowResponse)
	}
	return results
}
