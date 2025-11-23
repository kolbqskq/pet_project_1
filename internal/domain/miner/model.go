package miner

import (
	"MinersGame/internal/config"
	"MinersGame/pkg/errs"
	"sync"

	"github.com/google/uuid"
)

type Miner struct {
	ID         string
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
		ID:         uuid.New().String(),
		Config:     cfg,
		EnergyLeft: cfg.Energy,
	}, nil
}
