package http

import (
	"github.com/gin-gonic/gin"
	"github.com/idoceb00/elogap-api/internal/config"
	"github.com/idoceb00/elogap-api/internal/http/handlers"
	"github.com/idoceb00/elogap-api/internal/repository/memory"
	"github.com/idoceb00/elogap-api/internal/services"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	config.ApplyCORS(r)

	healthHandler, activityHandler, metricsHandler := buildHandlers()

	r.GET("/health", healthHandler.Get)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthHandler.Get)
		v1.GET("/activities", activityHandler.List)
		v1.GET("/activities/:id", activityHandler.GetByID)
		v1.GET("/metrics/summary", metricsHandler.Summary)
		v1.GET("/metrics/trends", metricsHandler.Trends)
	}

	return r
}

func buildHandlers() (*handlers.HealthHandler, *handlers.ActivityHandler, *handlers.MetricsHandler) {
	repo := memory.NewInMemoryActivityRepository()

	activitySvc := services.NewActivityService(repo)
	metricsSvc := services.NewMetricsService(repo)

	health := handlers.NewHealthHandler()
	activity := handlers.NewActivityHandler(activitySvc)
	metrics := handlers.NewMetricsHandler(metricsSvc)

	return health, activity, metrics
}
