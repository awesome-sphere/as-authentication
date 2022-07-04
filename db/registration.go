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

func CreateNewUser(username, hash, email string) {
	user := &models.User{
		Username:       username,
		HashedPassword: hash,
		Email:          email,
	}
	DB.Create(&user)
}
