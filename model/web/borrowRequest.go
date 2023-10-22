package web

import "time"

type BorrowCreateRequest struct {
	BookID uint      `json:"bookId" validate:"required"`
	UserID uint      `json:"userId" validate:"required"`
	Date   time.Time `json:"date"`
	Return time.Time `json:"return"`
	Status string    `json:"status" validate:"required"`
}

type BorrowUpdateRequest struct {
	BookID uint      `json:"bookId" validate:"required"`
	UserID uint      `json:"userId" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
	Return time.Time `json:"return" validate:"required"`
	Status string    `json:"status" validate:"required"`
}
