package domain

import "time"

type MetricsSummary struct {
	PlayerID           string  `json:"playerId"`
	Range              string  `json:"range"`
	TotalMatches       int     `json:"totalMatches"`
	WinRate            float64 `json:"winRate"` // 0..100
	AvgKDA             float64 `json:"avgKda"`
	AvgCsPerMin        float64 `json:"avgCsPerMin"`
	AvgVision          float64 `json:"avgVision"`
	AvgDamage          float64 `json:"avgDamage"`
	AvgDurationSeconds float64 `json:"avgDurationSeconds"`
}

type TrendPoint struct {
	Date     string  `json:"date"` // YYYY-MM-DD
	WinRate  float64 `json:"winRate"`
	KDA      float64 `json:"kda"`
	CsPerMin float64 `json:"csPerMin"`
}

type MetricsTrends struct {
	PlayerID string       `json:"playerId"`
	Range    string       `json:"range"`
	Points   []TrendPoint `json:"points"`
}

func DayKey(t time.Time) string {
	return t.UTC().Format("2006-01-02")
}
