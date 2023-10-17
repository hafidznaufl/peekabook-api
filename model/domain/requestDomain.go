package domain

import "time"

type Request struct {
	ID      uint
	Message string
	AdminID uint
	Admin   Admin
	UserID  uint
	User    User
	Date    time.Time
}
