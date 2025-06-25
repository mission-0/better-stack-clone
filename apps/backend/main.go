package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/controllers"
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
	router.GET("/health", controllers.HealthCheckup)
	router.POST("/signup", controllers.SignUpController)
	router.POST("/signin", controllers.SignInController)
	router.POST("/add", controllers.AddPostController)

	router.Run()
}
