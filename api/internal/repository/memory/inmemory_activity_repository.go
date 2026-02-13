package memory

import (
	"time"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository"
)

type InMemoryActivityRepository struct {
	data []domain.Activity
}

func NewInMemoryActivityRepository() repository.ActivityRepository {
	now := time.Now().UTC()

	return &InMemoryActivityRepository{
		data: []domain.Activity{
			{
				ID:        "match_1",
				PlayerID:  "player_1",
				Result:    domain.ResultWin,
				QueueType: "RANKED_SOLO",
				Champion:  "Ahri",
				Role:      "MID",
				Kills:     10,
				Deaths:    2,
				Assists:   8,
				CS:        210,
				Duration:  1714,
				Damage:    24500,
				Vision:    32,
				PlayedAt:  now.Add(-2 * time.Hour),
			},
			{
				ID:        "match_2",
				PlayerID:  "player_1",
				Result:    domain.ResultLoss,
				QueueType: "RANKED_SOLO",
				Champion:  "Jinx",
				Role:      "ADC",
				Kills:     4,
				Deaths:    7,
				Assists:   6,
				CS:        260,
				Duration:  1980,
				Damage:    30200,
				Vision:    18,
				PlayedAt:  now.Add(-6 * time.Hour),
			},
		},
	}
}

func (r *InMemoryActivityRepository) List() ([]domain.Activity, error) {
	// return a copy to avoid accidental mutation
	out := make([]domain.Activity, len(r.data))
	copy(out, r.data)
	return out, nil
}

func (r *InMemoryActivityRepository) FindByID(id string) (*domain.Activity, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			// return copy
			a := r.data[i]
			return &a, nil
		}
	}
	return nil, repository.ErrNotFound
}
