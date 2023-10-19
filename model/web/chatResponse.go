package web

import "time"

type ChatResponse struct {
	ID      uint      `json:"id"`
	Message string    `json:"message"`
	AdminID uint      `json:"adminId"`
	UserID  uint      `json:"userId"`
	Date    time.Time `json:"date"`
}
