package main

import (
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/jwt"
	"github.com/awesome-sphere/as-authentication/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// NOTE: Change to ReleaseMode when releasing the app
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// initialze database
	db.InitializeDatabase()
	jwt.InitializeJWTSettings()

	router.POST("/authen/login", service.Login)
	router.POST("/authen/signup", service.Signup)

	router.POST("/authen/update-hist", service.Booking)
	router.GET("/authen/get-hist", service.GetHistory)

	router.Run(":9001")
}
