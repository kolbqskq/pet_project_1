package saves

import "MinersGame/internal/domain/save"

type LoadManagerDeps struct {
	Name string
	GameLoadManager
}

func LoadGame(deps LoadManagerDeps) (*save.GameSave, error) {
	return deps.GameLoadManager.Load(deps.Name)
}
