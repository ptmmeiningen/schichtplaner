package models

import "time"

type Department struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	Name        string     `json:"name" gorm:"unique;not null"`
	Description string     `json:"description"`
	Color       string     `json:"color" gorm:"not null"`
	Users       []User     `json:"users" gorm:"many2many:user_departments;"`
}
