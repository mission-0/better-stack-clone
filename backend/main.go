package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/controllers"
	"github.com/mission-0/better-stack-backend/middlewares"
	utilities "github.com/mission-0/better-stack-backend/utilities"
)

func init() {
	fmt.Println("Init called")
	utilities.LoadEnvVaribales()
	utilities.ConnectToDb()
	// utilities.MigrateDB()
	// utilities.ConnectToRedis()
}

func main() {
	fmt.Println("Hello from go lang backend")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "UPDATE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	router.GET("/health", middlewares.Usermiddleware(), controllers.HealthCheckup)
	router.POST("/signup", controllers.SignUpController)
	router.POST("/signin", controllers.SignInController)
	router.GET("/mysites", middlewares.Usermiddleware(), controllers.AllSites)
	router.POST("/newsite", middlewares.Usermiddleware(), controllers.AddNewSiteController)
	router.GET("/logs", middlewares.Usermiddleware(), controllers.SiteLogs)

	router.Run()
}
