package saves

import (
	"MinersGame/internal/domain/equipment"
	"MinersGame/internal/domain/save"
	"MinersGame/internal/domain/stats"
)

type GameSaveManager interface {
	Save(save *save.GameSave) error
}

type BalanceManager interface {
	GetBallance() int
}

type StatsProvider interface {
	GetStats() ([]stats.MinerInfo, error)
}

type EquipmentManager interface {
	GetEquipmentsStatusOwn() map[string]equipment.Equipment
}

type GameLoadManager interface {
	Load(name string) (*save.GameSave, error)
}
