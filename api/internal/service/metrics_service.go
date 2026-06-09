package service

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

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

	var wins int
	var sumKDA, sumCsPerMin, sumVision, sumDamage, sumDuration float64

	for _, a := range items {
		if a.Result == domain.ResultWin {
			wins++
		}

		deaths := float64(a.Deaths)
		if deaths < 1 {
			deaths = 1
		}
		kda := (float64(a.Kills) + float64(a.Assists)) / deaths

		minutes := float64(a.Duration) / 60.0
		if minutes <= 0 {
			minutes = 1
		}
		csPerMin := float64(a.CS) / minutes

		sumKDA += kda
		sumCsPerMin += csPerMin
		sumVision += float64(a.Vision)
		sumDamage += float64(a.Damage)
		sumDuration += float64(a.Duration)

	}

	total := float64(len(items))
	summary.TotalMatches = len(items)
	summary.WinRate = (float64(wins) / total) * 100.0
	summary.AvgKDA = round2(sumKDA / total)
	summary.AvgCsPerMin = round2(sumCsPerMin / total)
	summary.AvgVision = round2(sumVision / total)
	summary.AvgDamage = round2(sumDamage / total)
	summary.AvgDurationSeconds = round2(sumDuration / total)

	return summary, nil
}

func (s *MetricsService) Trends(f MetricsFilter) (domain.MetricsTrends, error) {
	items, err := s.filteredActivities(f)
	if err != nil {
		return domain.MetricsTrends{}, err
	}

	out := domain.MetricsTrends{
		PlayerID: f.PlayerID,
		Range:    f.Range,
		Points:   []domain.TrendPoint{},
	}
	if len(items) == 0 {
		return out, nil
	}

	// agrupar por día
	type agg struct {
		total int
		wins  int
		sumK  float64
		sumC  float64
	}

	byDay := map[string]*agg{}

	for _, a := range items {
		key := domain.DayKey(a.PlayedAt)

		if _, ok := byDay[key]; !ok {
			byDay[key] = &agg{}
		}
		g := byDay[key]
		g.total++

		if a.Result == domain.ResultWin {
			g.wins++
		}

		deaths := float64(a.Deaths)
		if deaths < 1 {
			deaths = 1
		}
		kda := (float64(a.Kills) + float64(a.Assists)) / deaths

		minutes := float64(a.Duration) / 60.0
		if minutes <= 0 {
			minutes = 1
		}
		csPerMin := float64(a.CS) / minutes

		g.sumK += kda
		g.sumC += csPerMin
	}

	// ordenar por fecha ascendente
	keys := make([]string, 0, len(byDay))
	for k := range byDay {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		g := byDay[k]
		total := float64(g.total)
		p := domain.TrendPoint{
			Date:     k,
			WinRate:  round2((float64(g.wins) / total) * 100.0),
			KDA:      round2(g.sumK / total),
			CsPerMin: round2(g.sumC / total),
		}
		out.Points = append(out.Points, p)
	}

	return out, nil
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

	from, ok, err := rangeToFromUTC(f.Range)
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

func rangeToFromUTC(r string) (time.Time, bool, error) {
	switch strings.ToLower(strings.TrimSpace(r)) {
	case "", "30d":
		return time.Now().UTC().Add(-30 * 24 * time.Hour), true, nil
	case "7d":
		return time.Now().UTC().Add(-7 * 24 * time.Hour), true, nil
	case "90d":
		return time.Now().UTC().Add(-90 * 24 * time.Hour), true, nil
	case "season":
		// MVP: tratamos season como 90d por ahora
		return time.Now().UTC().Add(-90 * 24 * time.Hour), true, nil
	case "all":
		return time.Time{}, false, nil
	default:
		return time.Time{}, false, fmt.Errorf("invalid range (use 7d|30d|90d|season|all)")
	}
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}
