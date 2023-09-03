package models

type Photo struct {
	ID 			int 		`gorm:"primaryKey"`
	Title 		string 	`json:"title"`
	Caption 		string 	`json:"caption"`
	PhotoUrl 	string 	`json:"photoUrl"`
	UserID 		int
}

type UpdatePhoto struct {
	Title 		string 	`validate:"required" json:"title"`
	Caption 		string 	`validate:"required" json:"caption"`
	PhotoUrl 	string 	`validate:"required" json:"photoUrl"`
}