package db

import (
	"github.com/awesome-sphere/as-authentication/models"
	"gorm.io/gorm"
)

func GetUser(username string) (*gorm.DB, models.User) {
	var user models.User
	tx := DB.First(&user, "username = ?", username)
	return tx, user
}
