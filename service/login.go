package service

import (
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// TODO: Serialize inputs (username, password)
	var input serializer.LoginSerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username and Password are required.", // cast it to string before showing
		})
		return
	}

	user := db.GetUser(input.Username)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username does not exist.", // cast it to string before showing
		})
		return
	}

	// TODO: Check if password is correct
	match, err := argon2id.ComparePasswordAndHash(input.Password, user.HashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Match: %v", match)
	c.JSON(http.StatusFound, gin.H{
		"message": "Login successful.", // cast it to string before showing
	})
}
