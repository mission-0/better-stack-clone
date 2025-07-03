package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mission-0/better-stack-backend/dto"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddNewSiteController(ctx *gin.Context) {

	var input dto.WebsiteInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	validate := utilities.NewValidator()

	if err := validate.Struct(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  utilities.FormatValidationErrors(err),
		})
		return
	}

	registerNewSite := models.Website{
		Url:     input.Url,
		Regions: input.Regions,
		UserId:  input.UserId,
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
