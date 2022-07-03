package service

import (
	"errors"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {

	var user serializer.SignupSerializer
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(), // cast it to string before showing
		})
		return
	}

	username := user.Username
	password := user.Password
	email := user.Email

	if !db.IsValidEmail(email) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "This email has already been taken.",
		})
		return
	}

	tx, _ := db.GetUser(username)
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "This username has already been taken.",
		})
		return
	}

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	db.CreateNewUser(username, hash, email)
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Succeeded!",
	})
}
