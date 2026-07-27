package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/idoceb00/elogap-api/internal/services"
)

type MetricsHandler struct {
	svc *services.MetricsService
}

func NewMetricsHandler(svc *services.MetricsService) *MetricsHandler {
	return &MetricsHandler{svc: svc}
}

func (h *MetricsHandler) Summary(c *gin.Context) {
	playerID := strings.TrimSpace(c.Query("playerId"))
	rng := strings.TrimSpace(c.DefaultQuery("range", "30d"))

	res, err := h.svc.Summary(services.MetricsFilter{
		PlayerID: playerID,
		Range:    rng,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *MetricsHandler) Trends(c *gin.Context) {
	playerID := strings.TrimSpace(c.Query("playerId"))
	rng := strings.TrimSpace(c.DefaultQuery("range", "30d"))

	res, err := h.svc.Trends(services.MetricsFilter{
		PlayerID: playerID,
		Range:    rng,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
