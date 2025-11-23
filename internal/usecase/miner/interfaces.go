package miner

import (
	"MinersGame/internal/domain/miner"
	"context"
)

type MinerManager interface {
	Mine(ctx context.Context, m *miner.Miner)
}

type BallanceManager interface {
	SpendCoal(amount int) error
}
