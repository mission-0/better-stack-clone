package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AllSites(ctx *gin.Context) {
	var userID uuid.UUID
	var err error
	var websites []models.Website

	userIDInterface, isOk := ctx.Get("userId")
	if !isOk {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authenticated",
		})
	}

	// userId from the jwt might be a string to parsing it to uuid

	if userIDStr, ok := userIDInterface.(string); ok {
		userID, err = uuid.Parse(userIDStr)
		fmt.Println("userID", userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "UserId format is not valid",
				"err":     err,
			})
			return
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid User type",
		})
		return
	}

	res := utilities.DB.Where("user_id = ?", userID).Find(&websites)
	if res.Error != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Db query failes",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "db query acceptedd",
		"query":   websites,
	})
}
