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
	return string(bytes), err
}

func SignUpController(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
		return
	}

	validate := utilities.NewValidator()

	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation Failed",
			"errors":  utilities.FormatValidationErrors(err),
		})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		fmt.Println("something went wrong while hashing password ")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	registerNewUser := models.User{Email: user.Email, Password: hashedPassword, Fullname: user.Fullname}
	res := utilities.DB.Create(&registerNewUser)

	if res.Error != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "email might already exists",
			"error":   res.Error,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Signup sucessfull",
		"user":    registerNewUser,
	})
}
