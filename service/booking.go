package service

import (
	"encoding/json"
	"net/http"

	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/jwt"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
)

func Booking(c *gin.Context) {
	authorized, claims := jwt.AuthorizeToken(c)
	if !authorized {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized. Invalid or expired token.",
		})
		return
	}
	var input serializer.BookingSerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	booking, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong when converting to JSON.",
		})
		return
	}

	user_id, ok := claims["user_id"].(float64)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is not a number."})
	}
	db.UpdateBookingHistory(int64(user_id), booking)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully added booking to History!",
	})
}
