package service

import (
	"strings"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository"
)

type ActivityService struct {
	repo repository.ActivityRepository
}

type ListActivitiesFilter struct {
	Result   *domain.MatchResult
	Champion *string
}

func NewActivityService(repo repository.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) List(filter ListActivitiesFilter) ([]domain.Activity, error) {
	items, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	out := make([]domain.Activity, 0, len(items))
	for _, a := range items {
		if filter.Result != nil && a.Result != *filter.Result {
			continue
		}
		if filter.Champion != nil && !strings.EqualFold(a.Champion, *filter.Champion) {
			continue
		}
		out = append(out, a)
	}
	return out, nil
}

func (s *ActivityService) FindByID(id string) (*domain.Activity, error) {
	return s.repo.FindByID(id)
}
