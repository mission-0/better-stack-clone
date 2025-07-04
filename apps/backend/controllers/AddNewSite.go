package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddNewSiteController(ctx *gin.Context) {

	var newSite models.Website

	if err := ctx.ShouldBindJSON(&newSite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	validate := utilities.NewValidator()

	if err := validate.Struct(newSite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  utilities.FormatValidationErrors(err),
		})
		return
	}

	registerNewSite := models.Website{
		Url:     newSite.Url,
		Regions: newSite.Regions,
		UserId:  newSite.UserId,
	}

	res := utilities.DB.Create(&registerNewSite)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "foreign key") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid UserId: user does not exist",
				"error":   res.Error.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create website",
				"error":   res.Error.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Website added successfully",
	})
}
