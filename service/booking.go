package service

import (
	"encoding/json"
	"net/http"

	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/serializer"
	"github.com/gin-gonic/gin"
)

func Booking(c *gin.Context) {
	var input serializer.BookingSerializer
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		print(err.Error())
		return
	}
	booking, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong when converting to JSON.",
		})
		return
	}

	db.UpdateBookingHistory(input.UserID, booking)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully added booking to History!",
	})
}
