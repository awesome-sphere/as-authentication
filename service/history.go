package service

import (
	"net/http"

	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/db/models"
	"github.com/awesome-sphere/as-authentication/jwt"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) {
	authorized, claims := jwt.AuthorizeToken(c)
	if !authorized {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized. Invalid or expired token.",
		})
		return
	}

	user_id, ok := claims["user_id"].(float64)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is not a number."})
	}

	var booking []serializer.HistorySerializer

	if err := db.DB.Model(&models.Booking{}).Find(&booking, "user_id", user_id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"history": []byte{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history": booking,
	})
}
