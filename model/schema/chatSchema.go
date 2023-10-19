package schema

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Message   string         `json:"message"`
	AdminID   uint           `gorm:"index"`
	Admin     Admin          `gorm:"foreignKey:AdminID"`
	UserID    uint           `gorm:"index"`
	User      User           `gorm:"foreignKey:UserID"`
	Date      time.Time      `json:"date"`
}
