package domain

import "time"

type Chat struct {
	ID      uint
	Message string
	AdminID uint
	UserID  uint
	Date    time.Time
}
