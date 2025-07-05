package controllers

import "github.com/gin-gonic/gin"

func HealthCheckup(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Server is Healthy (100% health)",
	})
}
