package handler

import (
	"MinersGame/internal/domain/save"
	"MinersGame/internal/usecase/saves"
	"MinersGame/pkg/res"
	"net/http"
	"time"
)

type SaveHandler struct {
	saves.GameSaveManager
	saves.BalanceManager
	saves.EquipmentManager
	saves.StatsProvider
	saves.GameLoadManager
	LoadFunc  func(*save.GameSave)
	TimeStart *time.Time
}
type SaveHandlerDeps struct {
	saves.GameSaveManager
	saves.BalanceManager
	saves.EquipmentManager
	saves.StatsProvider
	saves.GameLoadManager
	LoadFunc  func(*save.GameSave)
	TimeStart *time.Time
}

func NewSaveHandler(router *http.ServeMux, deps SaveHandlerDeps) {
	handler := &SaveHandler{
		GameSaveManager:  deps.GameSaveManager,
		BalanceManager:   deps.BalanceManager,
		EquipmentManager: deps.EquipmentManager,
		StatsProvider:    deps.StatsProvider,
		GameLoadManager:  deps.GameLoadManager,
		LoadFunc:         deps.LoadFunc,
		TimeStart:        deps.TimeStart,
	}

	router.HandleFunc("GET /game/save", handler.Save())
	router.HandleFunc("GET /game/load", handler.Load())
}

func (handler *SaveHandler) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			res.Json(w, 400, "name is required")
			return
		}
		if err := saves.SaveGame(saves.SaveManagerDeps{
			Name:             name,
			GameStartedAt:    *handler.TimeStart,
			GameSaveManager:  handler.GameSaveManager,
			BalanceManager:   handler.BalanceManager,
			EquipmentManager: handler.EquipmentManager,
			StatsProvider:    handler.StatsProvider,
		}); err != nil {
			res.Json(w, 0, err)
			return
		}
		res.Json(w, 200, "game saved")
	}
}

func (handler *SaveHandler) Load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			res.Json(w, 400, "name is required")
			return
		}

		save, err := handler.GameLoadManager.Load(name)
		if err != nil {
			res.Json(w, 0, err)
			return
		}
		handler.LoadFunc(save)

		res.Json(w, 200, "game loaded")
	}
}
