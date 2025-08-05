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
	"github.com/robfig/cron"
)

type logResult struct {
	latency     string
	status      string
	err         string
	website     models.Website
	isWebsiteUp bool
}

var userID uuid.UUID
var err error

func AddSiteLogs() {
	go func() {

		fmt.Println("running")
		// var websiteLogs models.Logs
		var websites []models.Website

		fmt.Println("invoked")

		fmt.Println("Calling fns for logs...")
		response := utilities.DB.Preload("User").Where("user_id = ?", userID).Find(&websites)

		if response.Error != nil {
			log.Fatal("Find query fails with error ", response.Error)
		}

		/// channels and go routines

		resultOfChannels := make(chan logResult, len(websites))

		for _, website := range websites {
			go func(w models.Website) {

				fmt.Println("website before marshelling", website.User)

				latency, status, err, websiteStatus := pingsites.GetLatency(w.URL)

				resultOfChannels <- logResult{
					latency:     latency,
					status:      status,
					err:         err,
					website:     w,
					isWebsiteUp: websiteStatus,
				}

			}(website)
		}

		for i := 0; i < len(websites); i++ {
			resChan := <-resultOfChannels
			if resChan.err != "null" {
				fmt.Println("Error pinging site:", resChan.website.URL, resChan.err)

			}

			fmt.Println("Latency from fn", resChan.latency, "status:", resChan.status, "err", resChan.err, "websiteStatus", resChan.isWebsiteUp)
			newWebsiteLogs := models.Logs{
				Logs:        resChan.status,
				WebsiteID:   resChan.website.ID,
				Latency:     resChan.latency,
				Error:       resChan.err,
				Time:        time.Now(),
				IsWebsiteUp: resChan.isWebsiteUp,
			}
			res := utilities.DB.Create(&newWebsiteLogs)

			fmt.Println("All good till here")
			if res.Error != nil {
				if strings.Contains(res.Error.Error(), "foreign key") {

					fmt.Println("err", res.Error)
					return
				} else {
					fmt.Println("err", res.Error)

					return
				}
			}

		}

	}()
}

func SiteLogs(ctx *gin.Context) {
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

	c := cron.New()

	c.AddFunc("@every 00h00m3s", AddSiteLogs)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add cron job",
			"error":   err.Error(),
		})
		return
	}

	c.Start()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Job added successfully",
	})

}
