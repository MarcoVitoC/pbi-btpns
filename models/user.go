package models

import (
	"time"
)

type User struct {
	ID 				int 			`gorm:"primaryKey"`
	Username 		string 		`validate:"required" json:"username"`
	Email 			string 		`gorm:"unique" validate:"required,email" json:"email"`
	Password 		string 		`validate:"required,min=6" json:"password"`
	Photo Photo 					`gorm:"foreignKey:UserID"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type LoginRequest struct {
	Email 		string 	`validate:"required,email" json:"email"`
	Password 	string 	`validate:"required" json:"password"`
}

type UpdateUser struct {
	Username 	string 	`validate:"required" json:"username"`
	Email 		string 	`validate:"required,email" json:"email"`
	Password 	string 	`validate:"required" json:"password"`
}