package models

import "time"

type User struct {
	ID          uint         `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   *time.Time   `gorm:"index" json:"deleted_at"`
	FirstName   string       `json:"first_name" gorm:"not null"`
	LastName    string       `json:"last_name" gorm:"not null"`
	Email       string       `json:"email" gorm:"unique;not null"`
	Password    string       `json:"password" gorm:"not null"`
	Color       string       `json:"color" gorm:"not null"`
	IsAdmin     bool         `json:"is_admin" gorm:"default:false"`
	Departments []Department `json:"departments" gorm:"many2many:user_departments;"`
	Shifts      []Shift      `json:"shifts" gorm:"foreignKey:UserID"`
}
