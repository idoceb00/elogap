package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository"
	"github.com/idoceb00/elogap-api/internal/service"
)

type ActivityHandler struct {
	svc *service.ActivityService
}

func NewActivityHandler(svc *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{svc: svc}
}

func (h *ActivityHandler) List(c *gin.Context) {
	var filter service.ListActivitiesFilter

	if v := c.Query("result"); v != "" {
		r := domain.MatchResult(v)
		if r != domain.ResultWin && r != domain.ResultLoss {
			jsonError(c, http.StatusBadRequest, "invalid result (use win|loss)")
			return
		}
		filter.Result = &r
	}

	if v := c.Query("champion"); v != "" {
		filter.Champion = &v
	}

	items, err := h.svc.List(filter)
	if err != nil {
		jsonError(c, http.StatusInternalServerError, "failed to list activities")
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *ActivityHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	item, err := h.svc.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			jsonError(c, http.StatusNotFound, "activity not found")
			return
		}
		jsonError(c, http.StatusInternalServerError, "failed to get activity")
		return
	}

	c.JSON(http.StatusOK, item)
}

func jsonError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"error": msg})
}
