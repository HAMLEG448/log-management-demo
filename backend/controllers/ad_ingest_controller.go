package controllers

import (
	"net/http"

	"log-management/backend/config"
	"log-management/backend/entity"
	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func IngestADLog(c *gin.Context) {
	var input entity.ADLogInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	normalizedLog, err := sevices.NormalizeADLog(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to normalize AD log",
		})
		return
	}

	if err := config.DB.Create(&normalizedLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save AD log",
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
		"message": "AD log ingested successfully",
		"log":     normalizedLog,
	})
}
