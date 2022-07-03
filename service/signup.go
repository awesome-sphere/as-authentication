package service

import (
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/models"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	user := models.User{}
	c.Bind(&user)
	username := user.Username
	password := user.HashedPassword
	email := user.Email

	if !db.IsValidEmail(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "This email has already been used. Please register your account with different email address.",
		})
		return
	}

	if !db.IsValidUsername(username) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "This username has already been taken.",
		})
		return
	}

	if len(password) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Password is required!",
		})
		return
	} else if len(password) > 0 && len(password) <= 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Password must be longer than 8 characters!",
		})
		return
	}

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Error while encrypting!",
		})
		return
	}

	newAccount := &models.User{
		Username:       username,
		HashedPassword: hash,
		Email:          email,
	}

	db.CreateNewUser(*newAccount)
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Succeeded!",
	})
}
