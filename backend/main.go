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
	// utilities.MigrateDB()
	// utilities.ConnectToRedis()
}

func main() {
	fmt.Println("Hello from go lang backend")
	// currentTime, status, err := pingsites.GetLatency("https://jsonplaceholder.typicode.com/todos/")
	// if err != nil {
	// 	fmt.Println("error come")
	// }
	// fmt.Println("status: ", status)
	// fmt.Println("time is", currentTime)

	router := gin.Default()
	router.GET("/health", middlewares.Usermiddleware(), controllers.HealthCheckup)
	router.POST("/signup", controllers.SignUpController)
	router.POST("/signin", controllers.SignInController)
	router.GET("/mysites", middlewares.Usermiddleware(), controllers.AllSites)
	router.POST("/newsite", middlewares.Usermiddleware(), controllers.AddNewSiteController)
	router.GET("/logs", middlewares.Usermiddleware(), controllers.SiteLogs)

	router.Run()
}
