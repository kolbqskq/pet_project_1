package save

import (
	"MinersGame/internal/domain/equipment"
	"MinersGame/internal/domain/stats"
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GameSaveJSON struct {
	gorm.Model

	Name          string         `json:"name"`
	SaveAt        time.Time      `json:"saveAt"`
	GameStartedAt time.Time      `json:"gameStartedAt"`
	Wallet        int            `json:"wallet"`
	Equipments    datatypes.JSON `json:"equipments"`
	Miners        datatypes.JSON `json:"miners"`
}

type GameSave struct {
	Name          string                         `json:"name"`
	SaveAt        time.Time                      `json:"saveAt"`
	GameStartedAt time.Time                      `json:"gameStartedAt"`
	Wallet        int                            `json:"wallet"`
	Equipments    map[string]equipment.Equipment `json:"equipments"`
	Miners        []stats.MinerInfo              `json:"miners"`
}

func (g *GameSaveJSON) ToGameSave() (*GameSave, error) {
	var equipments map[string]equipment.Equipment
	if err := json.Unmarshal(g.Equipments, &equipments); err != nil {
		return nil, err
	}

	var miners []stats.MinerInfo
	if err := json.Unmarshal(g.Miners, &miners); err != nil {
		return nil, err
	}

	return &GameSave{
		Name:          g.Name,
		SaveAt:        g.SaveAt,
		GameStartedAt: g.GameStartedAt,
		Wallet:        g.Wallet,
		Equipments:    equipments,
		Miners:        miners,
	}, nil
}

func (g *GameSave) ToGameSaveJSON() (*GameSaveJSON, error) {
	equipments, err := json.Marshal(g.Equipments)
	if err != nil {
		return nil, err
	}

	miners, err := json.Marshal(g.Miners)
	if err != nil {
		return nil, err
	}

	return &GameSaveJSON{
		Name:          g.Name,
		SaveAt:        g.SaveAt,
		GameStartedAt: g.GameStartedAt,
		Wallet:        g.Wallet,
		Equipments:    equipments,
		Miners:        miners,
	}, nil
}
