package repository

import (
	"fmt"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository/record"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ActivityRepository interface {
	FindByPlayerID(puuid string) ([]domain.Activity, error)
	SaveMany(activities []domain.Activity) error
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db: db}
}

func (r *activityRepository) FindByPlayerID(puuid string) ([]domain.Activity, error) {
	var recs []record.Activity

	err := r.db.
		Where("player_id = ?", puuid).
		Order("played_at DESC").
		Find(&recs).Error
	if err != nil {
		return nil, fmt.Errorf("find activities by player id %q: %w", puuid, err)
	}

	activities := make([]domain.Activity, 0, len(recs))
	for _, rec := range recs {
		activities = append(activities, rec.ToDomain())
	}

	return activities, nil
}

func (r *activityRepository) SaveMany(activities []domain.Activity) error {
	if len(activities) == 0 {
		return nil
	}

	recs := make([]record.Activity, 0, len(activities))
	for _, a := range activities {
		recs = append(recs, record.ActivityFromDomain(a))
	}

	err := r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).
		Create(&recs).Error
	if err != nil {
		return fmt.Errorf("save activities: %w", err)
	}

	return nil
}
