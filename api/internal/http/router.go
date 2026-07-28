package http

import (
	"github.com/gin-gonic/gin"
	"github.com/idoceb00/elogap-api/internal/config"
	"github.com/idoceb00/elogap-api/internal/http/handlers"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	config.ApplyCORS(r)

	healthHandler := handlers.NewHealthHandler()

	r.GET("/health", healthHandler.Get)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthHandler.Get)
	}

	return r
}
