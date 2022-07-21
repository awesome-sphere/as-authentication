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
	authorized, _ := jwt.AuthorizeToken(c)
	if !authorized {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized. Invalid or expired token.",
		})
		return
	}
	var input serializer.HistorySerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var booking models.Booking

	if err := db.DB.First(&booking, "user_id", input.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history": booking,
	})
}
