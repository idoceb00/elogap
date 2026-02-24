package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository"
)

type MetricsService struct {
	repo repository.ActivityRepository
}

func NewMetricsService(repo repository.ActivityRepository) *MetricsService {
	return &MetricsService{repo: repo}
}

type MetricsFilter struct {
	PlayerID string
	Range    string
}

func (s *MetricsService) Summary(f MetricsFilter) (domain.MetricsSummary, error) {
	items, err := s.filteredActivities(f)
	if err != nil {
		return domain.MetricsSummary{}, err
	}

	summary := domain.MetricsSummary{
		PlayerID: f.PlayerID,
		Range:    f.Range,
	}

	if len(items) == 0 {
		return summary, nil
	}

	return summary, nil
}

func (s *MetricsService) filteredActivities(f MetricsFilter) ([]domain.Activity, error) {
	playerID := strings.TrimSpace(f.PlayerID)
	if playerID == "" {
		return nil, fmt.Errorf("playerId is required")
	}

	items, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	out := make([]domain.Activity, 0, len(items))
	for _, a := range items {
		if a.PlayerID != playerID {
			continue
		}
		if ok {
			if a.PlayedAt.UTC().Before(from) {
				continue
			}
		}
		out = append(out, a)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].PlayedAt.After(out[j].PlayedAt)
	})

	return out, nil
}
