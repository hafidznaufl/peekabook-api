package web

import (
	"time"
)

type ChatCreateRequest struct {
	Message string    `json:"message" validate:"required"`
	AdminID uint      `json:"adminId" validate:"required"`
	UserID  uint      `json:"userId" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
}

type ChatUpdateRequest struct {
	Message string    `json:"message" validate:"required"`
	AdminID uint      `json:"adminId" validate:"required"`
	UserID  uint      `json:"userId" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
}
