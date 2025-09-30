package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mission-0/better-stack-backend/models"
	"github.com/mission-0/better-stack-backend/utilities"
)

func AddNewSiteController(ctx *gin.Context) {
	var newSite models.Website
	var userID uuid.UUID
	var err error

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

	userIDInterface, isOk := ctx.Get("userId")
	if !isOk {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authenticated",
		})
	}

	// userId from the jwt might be a string to parsing it to uuid

	if userIDStr, ok := userIDInterface.(string); ok {
		userID, err = uuid.Parse(userIDStr)
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

	registerNewSite := models.Website{
		URL:          newSite.URL,
		Regions:      newSite.Regions,
		PingInterval: newSite.PingInterval,
		UserID:       userID,
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
