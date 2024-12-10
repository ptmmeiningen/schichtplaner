package models

import (
	"time"
)

type Todo struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
}
