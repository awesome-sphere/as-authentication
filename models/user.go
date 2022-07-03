package models

type User struct {
	Id             int    `json:"id" gorm:"primaryKey;autoincrement;not null"`
	Username       string `json:"username" gorm:"not null;unique"`
	HashedPassword string `json:"hashed_password" gorm:"not null"`
	Email          string `json:"email" gorm:"not null;unique"`
	JWTPrivateKey  string `json:"jwt_private_key"`
}
