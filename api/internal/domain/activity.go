package domain

import "time"

type MatchResult string

const (
	ResultWin  MatchResult = "win"
	ResultLoss MatchResult = "loss"
)

type Activity struct {
	ID        string      `json:"id"`
	PlayerID  string      `json:"playerId"`
	Result    MatchResult `json:"result"`
	QueueType string      `json:"queueType"`
	Champion  string      `json:"champion"`
	Role      string      `json:"role"`
	Kills     int         `json:"kills"`
	Deaths    int         `json:"deaths"`
	Assists   int         `json:"assists"`
	CS        int         `json:"cs"`
	Duration  int         `json:"duration"` // seconds
	Damage    int         `json:"damage"`
	Vision    int         `json:"vision"`
	PlayedAt  time.Time   `json:"playedAt"`
}
