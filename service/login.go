package service

import (
	"errors"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/jwt"
	"github.com/awesome-sphere/as-authentication/models"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func checkInput(c *gin.Context) (bool, bool, models.User) {
	var input serializer.LoginSerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username and Password are required.",
		})
		return true, false, models.User{}
	}

	tx, user := db.GetUser(input.Username)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username does not exist.",
		})
		return true, false, user
	}

	match, err := argon2id.ComparePasswordAndHash(input.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return true, false, user
	}
	return false, match, user
}

func Login(c *gin.Context) {
	// TODO: Serialize inputs (username, password)
	failed, match, user := checkInput(c)
	if failed {
		return
	}
	switch match {
	case true:
		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful!",
			"token":    jwt.GenerateJWT(user),
			"username": user.Username,
			"is_admin": user.IsAdmin,
			"email":    user.Email,
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is incorrect.",
		})
	}
}
