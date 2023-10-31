package web

import "time"

type BorrowResponse struct {
	ID        uint      `json:"id"`
	BookID    uint      `json:"bookId"`
	BookTitle string    `json:"bookTitle"`
	UserName  string    `json:"userName"`
	Date      time.Time `json:"date"`
	Return    time.Time `json:"return"`
	Status    string    `json:"status"`
}

type CreateBorrowResponse struct {
	ID     uint      `json:"id"`
	BookID uint      `json:"bookId"`
	UserID uint      `json:"userId"`
	Date   time.Time `json:"date"`
	Return time.Time `json:"return"`
	Status string    `json:"status"`
}
