package models

type User struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	JWTPrivateKey  string `json:"jwt_private_key"`
}
