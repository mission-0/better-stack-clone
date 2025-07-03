package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/dto"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddNewSiteController(ctx *gin.Context) {

	var input map[string]interface{}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	result := dto.WebSiteSchema.Parse(input, &models.Website{})

	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  result,
		})
	}

	var newSite models.Website
	newSite.Url = input["url"].(string)

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
