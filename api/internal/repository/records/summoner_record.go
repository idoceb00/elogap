package records

import (
	"time"

	"github.com/idoceb00/elogap-api/internal/domain"
)

type Summoner struct {
	PUUID     string    `gorm:"primaryKey;column:puuid"`
	GameName  string    `gorm:"column:game_name;not null"`
	TagLine   string    `gorm:"column:tag_line;not null"`
	Region    string    `gorm:"column:region;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Summoner) TableName() string {
	return "summoners"
}

func (r Summoner) ToDomain() domain.Summoner {
	return domain.Summoner{
		PUUID:     r.PUUID,
		GameName:  r.GameName,
		TagLine:   r.TagLine,
		Region:    domain.Region(r.Region),
		UpdatedAt: r.UpdatedAt,
	}
}

func SummonerFromDomain(s domain.Summoner) Summoner {
	return Summoner{
		PUUID:     s.PUUID,
		GameName:  s.GameName,
		TagLine:   s.TagLine,
		Region:    string(s.Region),
		UpdatedAt: s.UpdatedAt,
	}
}
