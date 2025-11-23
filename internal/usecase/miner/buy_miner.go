package miner

import (
	"MinersGame/internal/config"
	"MinersGame/internal/domain/miner"
	"context"
)

type BuyMinerDeps struct {
	MinerManager   MinerManager
	BalanceManager BallanceManager
	Class          string
	Ctx            context.Context
}

func BuyMiner(deps BuyMinerDeps) error {
	price := config.MinerPresets[deps.Class].Price
	if err := deps.BalanceManager.SpendCoal(int(price)); err != nil {
		return err
	}
	miner, err := miner.NewMiner(deps.Class)
	if err != nil {
		return err
	}
	deps.MinerManager.Mine(deps.Ctx, miner)
	return nil
}
