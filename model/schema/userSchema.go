package schema

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
