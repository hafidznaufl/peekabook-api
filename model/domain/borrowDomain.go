package domain

import "time"

type Borrow struct {
	ID        uint
	BookID    uint
	BookTitle string
	UserID    uint
	UserName  string
	Date      time.Time
	Return    time.Time
	Status    string
}
