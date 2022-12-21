package migration

import (
	"baharsah/helper/mysql"
	"baharsah/models"
	"log"
)

func RunMigration() {

	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Println(err)
		panic("Migrasi Gagal")
	} else {
		log.Println("Migrasi Berhasil")
	}
}
