package web

import "time"

type BorrowResponse struct {
	ID        uint      `json:"id"`
	BookTitle string    `json:"bookTitle"`
	UserName  string    `json:"userName"`
	Date      time.Time `json:"date"`
	Return    time.Time `json:"return"`
	Status    string    `json:"status"`
}
