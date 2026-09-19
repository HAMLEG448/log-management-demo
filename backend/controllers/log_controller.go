package controllers

import (
	"net/http"
	"time"

	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {
	tenant := c.GetString("effective_tenant")
	source := c.Query("source")
	eventType := c.Query("event_type")
	user := c.Query("user")
	search := c.Query("search")

	from := c.Query("from")
	to := c.Query("to")

	var fromTime *time.Time
	var toTime *time.Time

	if from != "" {
		parsedFrom, err := time.Parse(time.RFC3339, from)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid 'from' timestamp, use RFC3339 format",
			})
			return
		}

		fromTime = &parsedFrom
	}

	if to != "" {
		parsedTo, err := time.Parse(time.RFC3339, to)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid 'to' timestamp, use RFC3339 format",
			})
			return
		}

		toTime = &parsedTo
	}

	filter := sevices.LogFilter{
		Tenant:    tenant,
		Source:    source,
		EventType: eventType,
		User:      user,
		Search:    search,
		From:      fromTime,
		To:        toTime,
	}

	logs, err := sevices.GetLogs(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve logs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(logs),
		"logs":  logs,
	})
}
