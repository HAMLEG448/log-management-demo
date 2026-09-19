package controllers

import (
	"net/http"

	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func GetAlerts(c *gin.Context) {
	tenant := c.GetString("effective_tenant")

	alerts, err := sevices.GetAlerts(tenant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve alerts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(alerts),
		"alerts": alerts,
	})
}
