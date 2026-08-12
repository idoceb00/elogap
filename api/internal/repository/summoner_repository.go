package repository

import (
	"errors"
	"fmt"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository/record"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SummonerRepository interface {
	FindByRiotID(gameName, tagLine string) (domain.Summoner, error)
	Save(s domain.Summoner) error
}

type summonerRepository struct {
	db *gorm.DB
}

func NewSummonerRepository(db *gorm.DB) SummonerRepository {
	return &summonerRepository{db: db}
}

func (r *summonerRepository) FindByRiotID(gameName, tagLine string) (domain.Summoner, error) {
	var rec record.Summoner

	err := r.db.Where("game_name = ? AND tag_line = ?", gameName, tagLine).First(&rec).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Summoner{}, ErrNotFound
	} else if err != nil {
		return domain.Summoner{}, fmt.Errorf("find summoner by riot id %q#%q: %w", gameName, tagLine, err)
	}

	return rec.ToDomain(), nil
}

func (r *summonerRepository) Save(s domain.Summoner) error {
	rec := record.SummonerFromDomain(s)

	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "puuid"}},
		UpdateAll: true,
	}).Create(&rec).Error

	if err != nil {
		return fmt.Errorf("save summoner %q: %w", s.PUUID, err)
	}

	return nil
}
