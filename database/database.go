package database

import (
	"github.com/MarcoVitoC/pbi-btpns/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/pbi_btpns"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&app.User{})
	return db
}