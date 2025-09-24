package models

import (
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
