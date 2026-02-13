package repository

import "github.com/idoceb00/elogap-api/internal/domain"

type ActivityRepository interface {
	List() ([]domain.Activity, error)
	FindByID(id string) (*domain.Activity, error)
}
