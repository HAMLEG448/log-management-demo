package controllers

import (
	"net/http"

	"log-management/backend/entity"
	sevices "log-management/backend/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request entity.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, user, err := sevices.Login(
		request.Username,
		request.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,

		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
			"tenant":   user.Tenant,
		},
	})
}
