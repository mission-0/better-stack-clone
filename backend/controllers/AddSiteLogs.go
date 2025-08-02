package controllers

import (
	"encoding/json"
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

type logResult struct {
	latency string
	status  string
	err     string
	website models.Website
}

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
	response := utilities.DB.Preload("User").Where("user_id = ?", userID).Find(&websites)

	if response.Error != nil {
		log.Fatal("Find query fails with error ", response.Error)
	}

	/// channels and go routines

	resultOfChannels := make(chan logResult, len(websites))

	//serialising websites to valid json format before storing or Redis
	jsonSerialisation, jsonErr := json.Marshal(websites)
	if jsonErr != nil {
		fmt.Println("error serialising json")

	}

	for _, website := range websites {
		go func(w models.Website) {

			fmt.Println("website before marshelling", website.User)

			redisErr := utilities.RedisClient.Set(utilities.RedisContext, "websites"+"-"+w.User.Email, jsonSerialisation, 0)

			fmt.Println("rediserr:", redisErr)
			latency, status, err := pingsites.GetLatency(w.URL)
			resultOfChannels <- logResult{
				latency: latency,
				status:  status,
				err:     err,
				website: w,
			}

		}(website)
	}

	// var anyError bool

	for i := 0; i < len(websites); i++ {
		resChan := <-resultOfChannels
		if resChan.err != "null" {
			fmt.Println("Error pinging site:", resChan.website.URL, resChan.err)
			// anyError = true
			// continue
		}

		newWebsiteLogs := models.Logs{
			Logs:      resChan.status,
			WebsiteID: resChan.website.ID,
			Latency:   resChan.latency,
			Error:     resChan.err,
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

	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Site logs added",
	})
}
