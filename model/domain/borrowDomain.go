package domain

import "time"

type Borrow struct {
	ID     uint
	BookID uint
	Book   Book
	UserID uint
	User   User
	Date   time.Time
	Return time.Time
	Status string
}
