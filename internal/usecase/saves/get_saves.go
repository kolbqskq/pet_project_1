package saves

import "MinersGame/internal/domain/save"

type GetSavesDeps struct {
	GameSaveManager
}

func GetSaves(deps GetSavesDeps) ([]save.SaveInfo, error) {
	return deps.GetSave()
}
