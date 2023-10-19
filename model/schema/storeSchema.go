package schema

import (
	"time"
)

type Store struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	AdminID   uint      `gorm:"index"`
	Admin     Admin     `gorm:"foreignKey:AdminID"`
	BookID    uint      `gorm:"index"`
	Book      Book      `gorm:"foreignKey:BookID"`
	Date      time.Time `json:"date"`
}
