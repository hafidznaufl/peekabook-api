package web

import "time"

type BorrowCreateRequest struct {
	BookID uint      `json:"bookId" form:"bookId" validate:"required"`
	UserID uint      `json:"userId" form:"userId" validate:"required"`
	Date   time.Time `json:"date" form:"date" `
	Return time.Time `json:"return" form:"return"`
	Status string    `json:"status" form:"status" validate:"required"`
}

type BorrowUpdateRequest struct {
	BookID uint      `json:"bookId" form:"bookId" validate:"required"`
	UserID uint      `json:"userId" form:"userId" validate:"required"`
	Date   time.Time `json:"date" form:"date" `
	Return time.Time `json:"return" form:"return"`
	Status string    `json:"status" form:"status" validate:"required"`
}
