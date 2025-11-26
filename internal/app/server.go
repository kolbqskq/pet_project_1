package app

import (
	"MinersGame/internal/api/handler"
	"MinersGame/internal/domain/save"
	"MinersGame/pkg/res"
	"log/slog"
	"net/http"
)

func (app *App) RunServer(addr string) {
	router := http.NewServeMux()

	//Repositories:
	saveRepository := save.NewSaveRepository(app.Db)

	//Handlers:

	handler.NewMinersHandler(
		router,
		handler.MinersHandlerDeps{
			Ctx:             app.AppCtx,
			MinerManager:    app.MinerService,
			BallanceManager: app.WalletService,
		},
	)

	handler.NewEquipmentHandler(
		router,
		handler.EquipmentHandlerDeps{
			EquipmentManager: app.EquipmentService,
			BallanceManager:  app.WalletService,
		},
	)

	handler.NewStatsHandler(
		router,
		handler.StatsHandlerDeps{
			StatsProvider: app.StatsService,
		},
	)

	handler.NewEnterpriseHandler(
		router,
		handler.EnterpriseHandlerDeps{
			EnterpriseManager: app.EquipmentService,
			BallanceManager:   app.WalletService,
			StatsProvider:     app.StatsService,
			Ctx:               app.AppCtx,
			AppCancel:         app.AppCancel,
			TimeStart:         app.TimeStart,
		},
	)

	handler.NewSaveHandler(
		router, handler.SaveHandlerDeps{
			GameSaveManager:  saveRepository,
			BalanceManager:   app.WalletService,
			EquipmentManager: app.EquipmentService,
			StatsProvider:    app.StatsService,
			GameLoadManager:  saveRepository,
			LoadFunc:         app.LoadSave,
			TimeStart:        app.TimeStart,
		})

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		<-app.AppCtx.Done()
		if err := server.Shutdown(app.AppCtx); err != nil {
			slog.Error(err.Error())
		}
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error(err.Error())
	}

}

func (app *App) LoadHandler(router *http.ServeMux, repo *save.SaveRepository) {
	router.HandleFunc("GET /load", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			res.Json(w, 400, "name is required")
		}

	})
}
