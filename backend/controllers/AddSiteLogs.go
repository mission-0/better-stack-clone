package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mission-0/better-stack-backend/models"
	pingsites "github.com/mission-0/better-stack-backend/ping-sites"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddSiteLogs(ctx *gin.Context) {
	var websiteLogs models.Logs
	var websites []models.Website
	var userID uuid.UUID
	var err error

	fmt.Println("invoked")
	if err := ctx.ShouldBindJSON(&websiteLogs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("Binded json sucess....")

	userIDInterface, isOk := ctx.Get("userId")
	if !isOk {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authenticated",
		})
	}

	if userIDStr, ok := userIDInterface.(string); ok {
		userID, err = uuid.Parse(userIDStr)
		// fmt.Println("userID", userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "UserId format is not valid",
			})
			return
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid User type",
		})
		return
	}

	fmt.Println("Calling fns for logs...")
	response := utilities.DB.Where("user_id = ?", userID).Find(&websites)

	if response.Error != nil {
		log.Fatal("Find query fails with error ", response.Error)
	}

	for index := 0; index < len(websites); index++ {
		//calling fn for logs
		latency, status, err := pingsites.GetLatency(websites[index].URL)
		if err != nil {
			fmt.Println("error come")
		}
		// fmt.Println("status", status)

		fmt.Println("response from query url", websites[index].URL)
		fmt.Printf("type of latency is %T\n", latency)
		fmt.Println("latency", latency)

		newWebsiteLogs := models.Logs{
			Logs:      status,
			WebsiteID: websites[index].ID,
			Latency:   latency.String(),
			Time:      time.Now(),
		}
		res := utilities.DB.Create(&newWebsiteLogs)

		fmt.Println("All good till here")
		if res.Error != nil {
			if strings.Contains(res.Error.Error(), "foreign key") {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid UserId: user does not exist",
					"error":   res.Error.Error(),
				})
			} else {
				fmt.Println("err", res.Error)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to create website",
					"error":   res.Error.Error(),
				})
			}
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Site logs added",
		})
	}
}
