package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tensaitensai/TimeUS-api/pkg/model"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}
