package domain

import "time"

type Region string

const (
	RegionEUW Region = "euw1"
	RegionNA  Region = "na1"
	RegionKR  Region = "kr"
)

type Summoner struct {
	PUUID     string    `json:"puuid"`
	GameName  string    `json:"gameName"`
	TagLine   string    `json:"tagLine"`
	Region    Region    `json:"region"`
	UpdatedAt time.Time `json:"updatedAt"`
}
