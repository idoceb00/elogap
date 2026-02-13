package http

import (
	"github.com/gin-gonic/gin"
	"github.com/idoceb00/elogap-api/internal/config"
	"github.com/idoceb00/elogap-api/internal/http/handlers"
	"github.com/idoceb00/elogap-api/internal/repository/memory"
	"github.com/idoceb00/elogap-api/internal/service"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	config.ApplyCORS(r)

	healthHandler, activityHandler := buildHandlers()

	r.GET("/health", healthHandler.Get)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthHandler.Get)
		v1.GET("/activities", activityHandler.List)
		v1.GET("/activities/:id", activityHandler.GetByID)
	}

	return r
}

func buildHandlers() (*handlers.HealthHandler, *handlers.ActivityHandler) {
	repo := memory.NewInMemoryActivityRepository()
	svc := service.NewActivityService(repo)

	health := handlers.NewHealthHandler()
	activity := handlers.NewActivityHandler(svc)

	return health, activity
}
