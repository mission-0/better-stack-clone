package main

import (
	"fmt"

	pingsites "github.com/mission-0/better-stack-backend/ping-sites"
	utilities "github.com/mission-0/better-stack-backend/utilities"
)

func init() {
	fmt.Println("Init called")
	utilities.LoadEnvVaribales()
	utilities.ConnectToDb()
	// utilities.MigrateDB()
}

func main() {
	fmt.Println("Hello from go lang backend")
	currentTime, responseBody, err := pingsites.GetLatency("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		fmt.Println("error come")
	}
	fmt.Println("responseBody", responseBody)

	fmt.Println("time is", currentTime)
	/*
		router := gin.Default()
		router.GET("/health", middlewares.Usermiddleware(), controllers.HealthCheckup)
		router.POST("/signup", controllers.SignUpController)
		router.POST("/signin", controllers.SignInController)
		router.GET("/mysites", middlewares.Usermiddleware(), controllers.AllSites)
		router.POST("/newsite", middlewares.Usermiddleware(), controllers.AddNewSiteController)

		router.Run()
	*/
}
