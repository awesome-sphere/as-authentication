package main

import (
	"github.com/awesome-sphere/as-authentication/db"
	"github.com/awesome-sphere/as-authentication/service"
	"github.com/gin-gonic/gin"
)

func login() {

}

func main() {
	// NOTE: Change to ReleaseMode when releasing the app
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// initialze database
	db.Init()

	router.POST("/login", service.Login)
	router.POST("/login", service.Signup)
}
