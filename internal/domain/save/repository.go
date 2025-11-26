package save

import (
	"MinersGame/pkg/db"
	"MinersGame/pkg/errs"

	"github.com/gookit/slog"
)

type SaveRepository struct {
	Db *db.Db
}

func NewSaveRepository(db *db.Db) *SaveRepository {
	return &SaveRepository{Db: db}
}

func (repo *SaveRepository) Save(save *GameSave) error {
	saveJson, err := save.ToGameSaveJSON()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	result := repo.Db.Create(&saveJson)
	return result.Error
}

func (repo *SaveRepository) Load(name string) (*GameSave, error) {
	var saveJson GameSaveJSON
	result := repo.Db.Where("name = ?", name).First(&saveJson)
	if result.Error != nil {
		return nil, errs.NewSaveNotFound()
	}
	save, err := saveJson.ToGameSave()
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (repo *SaveRepository) GetSave() ([]SaveInfo, error) {
	var info []SaveInfo
	result := repo.Db.Model(&GameSaveJSON{}).
		Select("name", "save_at").
		First(&info)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return []SaveInfo{}, result.Error
	}
	return info, nil
}
