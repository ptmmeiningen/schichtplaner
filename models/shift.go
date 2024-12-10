package models

import (
	"time"
)

type Shift struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	StartTime   time.Time  `json:"start_time" gorm:"not null"`
	EndTime     time.Time  `json:"end_time" gorm:"not null"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	User        User       `json:"user"`
}
