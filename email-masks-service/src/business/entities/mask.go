package entities

import (
	"gorm.io/gorm"
	"time"
)

type EmailMask struct {
	gorm.Model
	Address     string    `json:"address" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	Enabled     bool      `json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
