package models

import "github.com/jackc/pgtype"

/*
Booking JSON Object
- movie_name: string
- user_name: string
- theater_name: string
- date: date
- time: time
- seat_number: string
*/

type Booking struct {
	ID      int64        `json:"id" gorm:"primary_key"`
	Booking pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null"`
	User    User         `gorm:"foreignkey:UserID"`
	UserID  int64        `json:"user_id" gorm:"not null"`
}
