package controllers

import (
	"net/http"
	"time"

	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	filter := sevices.DashboardFilter{
		Tenant: c.GetString("effective_tenant"),
		Source: c.Query("source"),
	}

	from := c.Query("from")

	if from != "" {
		parsedFrom, err := time.Parse(time.RFC3339, from)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid 'from' timestamp",
			})
			return
		}

		filter.From = &parsedFrom
	}

	to := c.Query("to")

	if to != "" {
		parsedTo, err := time.Parse(time.RFC3339, to)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid 'to' timestamp",
			})
			return
		}

		filter.To = &parsedTo
	}

	summary, err := sevices.GetDashboardSummary(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve dashboard data",
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}
