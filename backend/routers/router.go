package routers

import (
	"log-management/backend/controllers"
	"log-management/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	// Public
	router.GET("/health", controllers.HealthCheck)
	router.POST("/auth/login", controllers.Login)

	router.POST("/ingest", controllers.IngestLog)
	router.POST("/ingest/aws-file", controllers.IngestAWSFile)
	router.POST("/ingest/ad", controllers.IngestADLog)

	// Protected
	protected := router.Group("/")
	protected.Use(
		middleware.AuthMiddleware(),
		middleware.TenantMiddleware(),
	)

	protected.GET("/logs", controllers.GetLogs)
	protected.GET("/dashboard", controllers.GetDashboard)
	protected.GET("/alerts", controllers.GetAlerts)

	return router
}
