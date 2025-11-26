package saves

import (
	"MinersGame/internal/domain/save"
	"time"
)

type SaveManagerDeps struct {
	Name          string
	GameStartedAt time.Time
	GameSaveManager
	BalanceManager
	EquipmentManager
	StatsProvider
}

func SaveGame(deps SaveManagerDeps) error {
	miners, err := deps.StatsProvider.GetStats()
	if err != nil {
		return err
	}
	save := &save.GameSave{
		Name:          deps.Name,
		SaveAt:        time.Now(),
		GameStartedAt: deps.GameStartedAt,
		Wallet:        deps.BalanceManager.GetBallance(),
		Equipments:    deps.EquipmentManager.GetEquipmentsStatusOwn(),
		Miners:        miners,
	}

	deps.Save(save)

	return nil
}
