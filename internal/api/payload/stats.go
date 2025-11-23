package payload

import (
	"MinersGame/internal/domain/stats"
)

type GetStatsResponse struct {
	Miners []stats.MinerInfo `json:"miners"`
}
