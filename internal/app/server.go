package app

import (
	"MinersGame/internal/api/handler"
	"log/slog"
	"net/http"
)

func (app *App) RunServer(addr string) {
	router := http.NewServeMux()

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
			Cancel:            app.Cancel,
			TimeStart:         app.TimeStart,
		},
	)

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
