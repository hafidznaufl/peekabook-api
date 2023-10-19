package domain

import "time"

type Borrow struct {
	ID     uint
	BookID uint
	UserID uint
	Date   time.Time
	Return time.Time
	Status string
}
