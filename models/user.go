package models

type User struct {
	ID             int    `json:"id" gorm:"primaryKey;autoincrement;not null"`
	Username       string `json:"username" gorm:"not null;unique"`
	HashedPassword string `json:"hashed_password" gorm:"not null"`
	Email          string `json:"email" gorm:"not null;unique"`
	IsAdmin        bool   `json:"is_admin" gorm:"default:false"`
}
