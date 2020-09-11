package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tensaitensai/TimeUS-api/internal/model"
)

var db *gorm.DB
var count = 0

func init() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	name := os.Getenv("MYSQL_NAME")

	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name)
	connect(path)
	defer db.Close()

	if err := db.AutoMigrate(&model.User{}).Error; err != nil {
		panic(err.Error())
	}
	if err := db.AutoMigrate(&model.Post{}).Error; err != nil {
		panic(err.Error())
	}
}

func connect(path string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		log.Println("Not ready. Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			log.Fatal(err)
		}
		return connect(path)
	}
	log.Println("Successfully")
	return db, nil
}
