package mysql

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	errw := godotenv.Load(".env")
	if errw != nil {
		log.Fatal("Mysql : Error loading .env file")
	}

	var err error
	dsn := os.Getenv("MYSQL_DSN_CONNECTION_DETAILS")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("disini")
		panic(err)
	}

	log.Println("Connected to Databases")

}
