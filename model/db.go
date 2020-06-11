package model

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}
