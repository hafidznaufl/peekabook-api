package schema

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	AdminID   uint           `gorm:"index"`
	Admin     Admin          `gorm:"foreignKey:AdminID"`
	BookID    uint           `gorm:"index"`
	Book      Book           `gorm:"foreignKey:BookID"`
	Date      time.Time      `json:"date"`
}
