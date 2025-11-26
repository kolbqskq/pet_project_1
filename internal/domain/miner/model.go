package miner

import (
	"MinersGame/internal/config"
	"MinersGame/pkg/errs"
	"sync"
)

type Miner struct {
	Config     config.MinerConfig
	EnergyLeft int
	mu         sync.RWMutex
}

func NewMiner(class string) (*Miner, error) {
	cfg, ok := config.MinerPresets[class]
	if !ok {
		return &Miner{}, errs.NewInvalidClass() // error
	}
	return &Miner{
		Config:     cfg,
		EnergyLeft: cfg.Energy,
	}, nil
}
