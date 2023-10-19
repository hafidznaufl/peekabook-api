package schema

import (
	"time"
)

type Chat struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Message   string    `json:"message"`
	AdminID   uint      `gorm:"index"`
	Admin     Admin     `gorm:"foreignKey:AdminID"`
	UserID    uint      `gorm:"index"`
	User      User      `gorm:"foreignKey:UserID"`
	Date      time.Time `json:"date"`
}
