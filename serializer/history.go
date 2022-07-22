package serializer

import "gorm.io/datatypes"

type HistorySerializer struct {
	Booking datatypes.JSON `json:"booking" binding:"required"`
}
