package stats

import "MinersGame/internal/domain/miner"

type StatsMiners struct {
	StatMiners map[string]map[string]*miner.Miner
}

type MinerInfo struct {
	Class      string
	EnergyLeft int
}

type CountsMiners struct {
	Class string
	Count int
}
