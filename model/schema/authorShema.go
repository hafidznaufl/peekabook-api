package schema

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Name      string    `json:"name"`
	Books     []Book    `gorm:"foreignKey:AuthorID"`
}
