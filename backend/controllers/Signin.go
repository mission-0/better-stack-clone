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
	var loginRequest models.User

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
		return
	}

	var storedUser models.User
	result := utilities.DB.Where("email = ?", loginRequest.Email).First(&storedUser)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})
		return
	}

	if !checkUserPasswordWithHash(storedUser.Password, loginRequest.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect password",
		})
		return
	}

	jwtToken, err := createToken(storedUser.ID)
	if err != nil {
		fmt.Println("JWt err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create token",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signin successful",
		"JWT":     jwtToken,
	})
}
