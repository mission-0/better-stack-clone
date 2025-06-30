package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(simplePassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(simplePassword), 14)
	// note that the 14 written in the above GenerateFromPassword function is nothing but salts round
	return string(bytes), err
}

func SignUpController(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println("Json Bind Failed")
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Json Bind Failed",
		})
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		fmt.Println("something went wrong while hashing password ")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	registerNewUser := models.User{Name: user.Name, Email: user.Email, Password: hashedPassword, FullName: user.FullName}
	res := utilities.DB.Create(&registerNewUser)

	if res.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "email might already exists",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Signup sucessfull",
		"user":    registerNewUser,
	})
}
