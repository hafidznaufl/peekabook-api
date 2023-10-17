package domain

import "time"

type Store struct {
	ID      uint
	AdminID uint
	Admin   Admin
	BookID  uint
	Book    Book
	Date    time.Time
}
