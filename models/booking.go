package models

import "gorm.io/datatypes"

type Booking struct {
	ID      int64          `json:"id" gorm:"primaryKey;autoincrement;not null"`
	Booking datatypes.JSON `gorm:"type:jsonb;default:'[]';not null"`
	User    User           `gorm:"foreignkey:UserID"`
	UserID  int64          `json:"user_id" gorm:"not null"`
}
