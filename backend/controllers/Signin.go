package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
	"golang.org/x/crypto/bcrypt"
)

func checkUserPasswordWithHash(hashedPassword, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	return err == nil
}

func createToken(userId uuid.UUID) (string, error) {
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":  userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := unsignedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SignInController(ctx *gin.Context) {
	var user models.User

	ctx.ShouldBindJSON(&user)

	newUser := models.User{
		ID: user.ID,
	}

	fmt.Println("newUser", newUser)
	result := utilities.DB.First(&newUser)
	fmt.Println("result", result)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting user",
		})
		return
	}

	isCorrectPassword := checkUserPasswordWithHash(newUser.Password, user.Password)

	if !isCorrectPassword {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Incorrect Password",
		})
		return
	}

	jwtToken, err := createToken(newUser.ID)
	if err != nil {
		fmt.Println("JWt err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "Bad jwt call",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": " Signin route",
		"JWT":     jwtToken,
	})
}
