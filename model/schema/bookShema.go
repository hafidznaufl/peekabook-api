package schema

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Title     string `json:"title"`
	AuthorID  uint   `gorm:"index"`
	Page      int    `json:"page"`
	Years     int    `json:"years"`
	Publisher string `json:"publisher"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}
