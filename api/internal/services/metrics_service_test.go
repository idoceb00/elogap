package services_test

import (
	"testing"
	"time"

	"github.com/idoceb00/elogap-api/internal/domain"
	"github.com/idoceb00/elogap-api/internal/repository"
	"github.com/idoceb00/elogap-api/internal/services"
)

//
// ---- Fake Repository ----
//

type fakeActivityRepo struct {
	items []domain.Activity
}

func (f fakeActivityRepo) List() ([]domain.Activity, error) {
	out := make([]domain.Activity, len(f.items))
	copy(out, f.items)
	return out, nil
}

func (f fakeActivityRepo) FindByID(id string) (*domain.Activity, error) {
	for i := range f.items {
		if f.items[i].ID == id {
			a := f.items[i]
			return &a, nil
		}
	}
	return nil, repository.ErrNotFound
}

//
// ---- Tests ----
//

func TestMetricsService_Summary(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name    string
		repo    repository.ActivityRepository
		filter  services.MetricsFilter
		want    domain.MetricsSummary
		wantErr bool
	}{
		{
			name: "calculate summary for player with 2 wins and 1 loss",
			repo: fakeActivityRepo{
				items: []domain.Activity{
					{ID: "1", PlayerID: "p1", Result: domain.ResultWin, PlayedAt: now},
					{ID: "2", PlayerID: "p1", Result: domain.ResultLoss, PlayedAt: now},
					{ID: "3", PlayerID: "p1", Result: domain.ResultWin, PlayedAt: now},
					{ID: "4", PlayerID: "p2", Result: domain.ResultWin, PlayedAt: now},
				},
			},
			filter: services.MetricsFilter{
				PlayerID: "p1",
				Range:    "30d",
			},
			want: domain.MetricsSummary{
				TotalMatches: 3,
				WinRate:      66.67, // si redondeas a 2 decimales
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := services.NewMetricsService(tt.repo)

			got, err := s.Summary(tt.filter)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("Summary() unexpected error: %v", err)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("Summary() succeeded unexpectedly")
			}

			if got.TotalMatches != tt.want.TotalMatches {
				t.Errorf("TotalMatches = %d, want %d",
					got.TotalMatches, tt.want.TotalMatches)
			}

			// Comparación con tolerancia por floats
			if diff := got.WinRate - tt.want.WinRate; diff > 0.01 || diff < -0.01 {
				t.Errorf("WinRate = %.2f, want %.2f",
					got.WinRate, tt.want.WinRate)
			}
		})
	}
}

func TestMetricsService_Trends(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name    string
		repo    repository.ActivityRepository
		filter  services.MetricsFilter
		wantLen int
		wantErr bool
	}{
		{
			name: "generate trends for player grouped by date",
			repo: fakeActivityRepo{
				items: []domain.Activity{
					// Día 1
					{
						ID:       "1",
						PlayerID: "p1",
						Result:   domain.ResultWin,
						Kills:    5,
						Deaths:   1,
						Assists:  4,
						PlayedAt: now.Add(-24 * time.Hour),
					},
					{
						ID:       "2",
						PlayerID: "p1",
						Result:   domain.ResultLoss,
						Kills:    2,
						Deaths:   5,
						Assists:  1,
						PlayedAt: now.Add(-24 * time.Hour),
					},
					// Día 2
					{
						ID:       "3",
						PlayerID: "p1",
						Result:   domain.ResultWin,
						Kills:    10,
						Deaths:   2,
						Assists:  3,
						PlayedAt: now,
					},
					// Otro jugador
					{
						ID:       "4",
						PlayerID: "p2",
						Result:   domain.ResultWin,
						PlayedAt: now,
					},
				},
			},
			filter: services.MetricsFilter{
				PlayerID: "p1",
				Range:    "30d",
			},
			wantLen: 2, // esperamos 2 días distintos
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := services.NewMetricsService(tt.repo)

			got, err := s.Trends(tt.filter)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("Trends() unexpected error: %v", err)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("Trends() succeeded unexpectedly")
			}

			if len(got.Points) != tt.wantLen {
				t.Errorf("expected %d trend points, got %d",
					tt.wantLen, len(got.Points))
			}

			// Validación básica de contenido
			for _, p := range got.Points {
				if p.WinRate < 0 || p.WinRate > 100 {
					t.Errorf("invalid winRate value: %v", p.WinRate)
				}
				if p.KDA < 0 {
					t.Errorf("invalid KDA value: %v", p.KDA)
				}
			}
		})
	}
}
