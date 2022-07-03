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

func Login(c *gin.Context) {
	// TODO: Serialize inputs (username, password)
	var input serializer.LoginSerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username and Password are required.",
		})
		return
	}

	tx, user := db.GetUser(input.Username)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username does not exist.",
		})
		return
	}

	// TODO: Check if password is correct
	match, err := argon2id.ComparePasswordAndHash(input.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// TODO: JWT Token stuffs
	switch match {
	case true:
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful!",
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is incorrect.",
		})
	}
}
