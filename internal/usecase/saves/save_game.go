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
		Wallet:        deps.BalanceManager.GetBalance(),
		Equipments:    deps.EquipmentManager.GetEquipmentsStatusOwn(),
		Miners:        miners,
	}

	if err := deps.Save(save); err != nil {
		return err
	}

	return nil
}
