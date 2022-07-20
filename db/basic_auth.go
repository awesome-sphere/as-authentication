package db

import (
	"errors"

	"github.com/awesome-sphere/as-authentication/db/models"
	"gorm.io/gorm"
)

func GetUser(username string) (*gorm.DB, models.User) {
	var user models.User
	tx := DB.First(&user, "username = ?", username)
	return tx, user
}

func IsValidEmail(email string) bool {
	var user models.User
	isUsed := DB.Where("email = ?", email).First(&user)
	return errors.Is(isUsed.Error, gorm.ErrRecordNotFound)
}

func CreateNewUser(username, hash, email string) {
	user := &models.User{
		Username:       username,
		HashedPassword: hash,
		Email:          email,
	}
	DB.Create(&user)
}
