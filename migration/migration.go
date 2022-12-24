package migration

import (
	"baharsah/helper/mysql"
	"baharsah/models"

	log "github.com/sirupsen/logrus"
)

func RunMigration() {

	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Println(err)
		panic("Migrasi User Gagal")
	} else {
		log.Println("Migrasi User Berhasil")
	}

	err2 := mysql.DB.AutoMigrate(&models.Trips{})

	if err2 != nil {
		log.Println(err2)
		panic("Migrasi Trip Gagal")
	} else {
		log.Println("Migrasi Trip Berhasil")
	}

	err3 := mysql.DB.AutoMigrate(&models.ImageTrips{})

	if err3 != nil {
		log.Println(err3)
		panic("Migrasi ImageTrip Gagal")
	} else {
		log.Println("Migrasi ImageTrip Berhasil")
	}
	err4 := mysql.DB.AutoMigrate(&models.Countries{})

	if err4 != nil {
		log.Println(err4)
		panic("Migrasi ImageTrip Gagal")
	} else {
		log.Println("Migrasi ImageTrip Berhasil")
	}

	err5 := mysql.DB.AutoMigrate(&models.Transactions{})

	if err5 != nil {
		log.Println(err5)
		panic("Migrasi ImageTrip Gagal")
	} else {
		log.Println("Migrasi ImageTrip Berhasil")
	}
}
