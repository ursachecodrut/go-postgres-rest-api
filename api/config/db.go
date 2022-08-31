package config

import (
	"log"

	"github.com/ursachecodrut/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
