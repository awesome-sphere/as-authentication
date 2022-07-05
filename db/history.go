package db

import (
	"github.com/awesome-sphere/as-authentication/models"
	"gorm.io/datatypes"
)

func UpdateBookingHistory(user_id int64, booking datatypes.JSON) {
	DB.Create(&models.Booking{UserID: user_id, Booking: booking})
}
