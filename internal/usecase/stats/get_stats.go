package stats

import (
	"MinersGame/internal/config"
	"MinersGame/internal/domain/stats"
	"MinersGame/pkg/errs"
)

type StatsProviderDeps struct {
	StatsProvider StatsProvider
	Class         string
}

func GetStats(deps StatsProviderDeps) ([]stats.MinerInfo, error) {
	stats, err := deps.StatsProvider.GetStats()
	return stats, err
}

func GetStatsByClass(deps StatsProviderDeps) ([]stats.MinerInfo, error) {
	if _, ok := config.MinerPresets[deps.Class]; !ok {
		return nil, errs.NewInvalidClass()
	}
	stats, err := deps.StatsProvider.GetStatsByClass(deps.Class)
	return stats, err
}

func GetCounts(deps StatsProviderDeps) []stats.CountsMiners {
	res := deps.StatsProvider.GetCountsAllClass()
	return res
}
