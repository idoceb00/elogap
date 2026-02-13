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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config.ApplyCORS(r)

	repo := memory.NewInMemoryActivityRepository()
	svc := service.NewActivityService(repo)

	health := handlers.NewHealthHandler()
	activity := handlers.NewActivityHandler(svc)

	r.GET("/health", health.Get)

	v1 := r.Group("/v1")
	{
		v1.GET("/activities", activity.List)
		v1.GET("/activities/:id", activity.GetByID)
	}

	return r
}
