package save

import (
	"MinersGame/pkg/db"
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
		return err
	}

	result := repo.Db.Create(&saveJson)
	return result.Error
}

func (repo *SaveRepository) Load(name string) (*GameSave, error) {
	var saveJson GameSaveJSON
	result := repo.Db.Where("name = ?", name).First(&saveJson)
	if result.Error != nil {
		return nil, result.Error
	}
	save, err := saveJson.ToGameSave()
	if err != nil {
		return nil, err
	}
	return save, nil
}
