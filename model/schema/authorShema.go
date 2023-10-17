package schema

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int            `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	Books     []Book         `gorm:"foreignKey:AuthorID"`
}
