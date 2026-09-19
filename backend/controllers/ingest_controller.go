package controllers

import (
	"net/http"

	"log-management/backend/config"
	"log-management/backend/entity"
	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func IngestLog(c *gin.Context) {

	var request entity.IngestRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	normalizedLog := sevices.NormalizeLog(request)

	if err := config.DB.Create(&normalizedLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save log",
		})
		return
	}

	if err := sevices.CheckAlertRules(normalizedLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to check alert rules",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "log ingested successfully",
		"log":     normalizedLog,
	})
}
