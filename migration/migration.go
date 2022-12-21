package migration

import (
	"baharsah/models"
	"baharsah/pkg/mysql"
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
