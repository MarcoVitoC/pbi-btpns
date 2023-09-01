package app

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID int `gorm:"primaryKey"`
	Username string `valid:"-"`
	Email string `gorm:"unique" valid:"email"`
	Password string `valid:"minstringlength(6)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}