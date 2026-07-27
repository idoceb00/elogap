package records

import (
	"time"

	"github.com/idoceb00/elogap-api/internal/domain"
)

type Activity struct {
	ID        string    `gorm:"primaryKey;column:id"`
	PlayerID  string    `gorm:"column:player_id;not null;index"`
	Result    string    `gorm:"column:result;not null"`
	QueueType string    `gorm:"column:queue_type;not null"`
	Champion  string    `gorm:"column:champion;not null"`
	Role      string    `gorm:"column:role;not null"`
	Kills     int       `gorm:"column:kills;not null"`
	Deaths    int       `gorm:"column:deaths;not null"`
	Assists   int       `gorm:"column:assists;not null"`
	CS        int       `gorm:"column:cs;not null"`
	Duration  int       `gorm:"column:duration;not null"`
	Damage    int       `gorm:"column:damage;not null"`
	Vision    int       `gorm:"column:vision;not null"`
	PlayedAt  time.Time `gorm:"column:played_at;not null"`
}

func (Activity) TableName() string {
	return "activities"
}

func (r Activity) ToDomain() domain.Activity {
	return domain.Activity{
		ID:        r.ID,
		PlayerID:  r.PlayerID,
		Result:    domain.MatchResult(r.Result),
		QueueType: r.QueueType,
		Champion:  r.Champion,
		Role:      r.Role,
		Kills:     r.Kills,
		Deaths:    r.Deaths,
		Assists:   r.Assists,
		CS:        r.CS,
		Duration:  r.Duration,
		Damage:    r.Damage,
		Vision:    r.Vision,
		PlayedAt:  r.PlayedAt,
	}
}

func ActivityFromDomain(s domain.Activity) Activity {
	return Activity{
		ID:        s.ID,
		PlayerID:  s.PlayerID,
		Result:    string(s.Result),
		QueueType: s.QueueType,
		Champion:  s.Champion,
		Role:      s.Role,
		Kills:     s.Kills,
		Deaths:    s.Deaths,
		Assists:   s.Assists,
		CS:        s.CS,
		Duration:  s.Duration,
		Damage:    s.Damage,
		Vision:    s.Vision,
		PlayedAt:  s.PlayedAt,
	}
}
