package db

import (
	"log"

	"github.com/awesome-sphere/as-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dbURL := "postgres://pkinwza:securepassword@localhost:5432/as-user"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Booking{}, &models.User{})

	DB = db
}
