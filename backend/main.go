package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/controllers"
	"github.com/mission-0/better-stack-backend/middlewares"
	utilities "github.com/mission-0/better-stack-backend/utilities"
)

func init() {
	fmt.Println("Init called")
	utilities.LoadEnvVaribales()
	utilities.ConnectToDb()
	utilities.MigrateDB()
}

func main() {
	fmt.Println("Hello from go lang backend")
	router := gin.Default()
	router.GET("/health", middlewares.Usermiddleware(), controllers.HealthCheckup)
	router.POST("/signup", controllers.SignUpController)
	router.POST("/signin", controllers.SignInController)
	router.GET("/mysites", middlewares.Usermiddleware(), controllers.AllSites)
	router.POST("/newsite", middlewares.Usermiddleware(), controllers.AddNewSiteController)

	router.Run()
}
