package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/dto"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(simplePassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(simplePassword), 14)
	return string(bytes), err
}

func SignUpController(ctx *gin.Context) {

	var input map[string]interface{}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	result := dto.SignupSchema.Parse(input, &models.User{})
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  result,
		})
		return
	}

	var user models.User

	user.Name = input["name"].(string)
	user.Email = input["email"].(string)
	user.Password = input["password"].(string)
	if Fullname, ok := input["Fullname"].(string); ok {
		user.Fullname = Fullname
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		fmt.Println("something went wrong while hashing password ")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	registerNewUser := models.User{Name: user.Name, Email: user.Email, Password: hashedPassword, Fullname: user.Fullname}
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
