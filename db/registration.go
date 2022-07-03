package db

import (
	"errors"

	"github.com/awesome-sphere/as-authentication/models"
	"gorm.io/gorm"
)

func IsValidEmail(email string) bool {
	var user models.User
	isUsed := DB.Where("email = ?", email).First(&user)
	return errors.Is(isUsed.Error, gorm.ErrRecordNotFound)
}

func IsValidUsername(username string) bool {
	var user models.User
	isUsed := DB.Where("username = ?", username).First(&user)
	return errors.Is(isUsed.Error, gorm.ErrRecordNotFound)
}

func CreateNewUser(user models.User) {
	DB.Create(&user)
}
