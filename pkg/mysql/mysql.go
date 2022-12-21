package mysql

import (
	"log"

	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {

	var err error
	dsn := "root:testserver@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("disini")
		panic(err)
	}

	log.Println("Connected to Databases")

}
