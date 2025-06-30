package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddNewSiteController(ctx *gin.Context) {
	var newSite models.Website
	if err := ctx.ShouldBindJSON(&newSite); err != nil {
		fmt.Println("Json Bind Failed")
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Json Bind Failed",
		})
	}

	ok := models.IsValidRegion(newSite.Regions)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "please enter a valid Region",
		})
		return
	}

	registerNewSite := models.Website{Regions: newSite.Regions, UserId: newSite.UserId, Url: newSite.Url}

	res := utilities.DB.Create(&registerNewSite)

	if res.Error != nil {
		fmt.Println("Error inserting idea")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messagee": "Error inserting data to Website field in db ",
			"err":      res.Error,
		})
		return
	}

	ctx.JSON((http.StatusAccepted), gin.H{
		"messagee": "dataInserted",
	})
}
