package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		userTenant := c.GetString("tenant")

		if role == "viewer" {
			if userTenant == "" {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "viewer has no tenant assigned",
				})
				c.Abort()
				return
			}

			// Viewer ถูกบังคับให้ใช้ tenant ของตัวเองเสมอ
			c.Set("effective_tenant", userTenant)

			c.Next()
			return
		}

		// Admin สามารถเลือก tenant จาก query ได้
		c.Set("effective_tenant", c.Query("tenant"))

		c.Next()
	}
}
