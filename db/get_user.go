package db

import "github.com/awesome-sphere/as-authentication/models"

func GetUser(username string) models.User {
	var user models.User
	DB.Where("username = ?", username).First(&user)
	return user
}
