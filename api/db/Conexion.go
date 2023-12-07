package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DNS = "root:@tcp(127.0.0.1:3306)/citas?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func DBConection() {
	var error error

	DB, error = gorm.Open(mysql.New(mysql.Config{
		DSN: DNS,
	}),				
	&gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to database!")
	} else {
		log.Println("Connection to database established")
	}
}