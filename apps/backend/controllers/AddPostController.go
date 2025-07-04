package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func init() {
	utilities.ConnectToDb()
}
func AddPostController(ctx *gin.Context) {
	user := models.User{Email: "hii@hello.com", Password: "test", Fullname: "shishu"}

	fmt.Println("user ", &user)
	result := utilities.DB.Create(&user)

	if result.Error != nil {
		fmt.Println("can't create user in DB")
		return

	}

	ctx.JSON(200, gin.H{
		"Usrr": user,
	})
}
