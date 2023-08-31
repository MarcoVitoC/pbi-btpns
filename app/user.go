package app

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	id int `gorm:"primaryKey" valid:"required"`
	username string `valid:"required"`
	email string `valid:"required"`
	password string `valid:"required,length(6)"`
}