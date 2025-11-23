package stats

import "MinersGame/internal/domain/stats"

type StatsProvider interface {
	GetStats() ([]stats.MinerInfo, error)
	GetStatsByClass(class string) ([]stats.MinerInfo, error)
	GetCountsAllClass() []stats.CountsMiners
}
