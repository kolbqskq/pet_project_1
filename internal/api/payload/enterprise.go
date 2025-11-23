package payload

import (
	"MinersGame/internal/domain/stats"
)

type EndGameResponse struct {
	Balance      int                  `json:"balance"`
	Miners       []stats.MinerInfo    `json:"miners"`
	CountMiners  []stats.CountsMiners `json:"countMiners"`
	GameDuration string               `json:"gameDuration"`
}

type GetBalanceResponse struct {
	Coal int
}
