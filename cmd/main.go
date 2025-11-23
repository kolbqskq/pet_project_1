package main

import (
	"MinersGame/internal/api/handler"
	"MinersGame/internal/domain/equipment"
	"MinersGame/internal/domain/miner"
	"MinersGame/internal/domain/stats"
	"MinersGame/internal/domain/wallet"
	"MinersGame/internal/types"
	"MinersGame/pkg/event"
	"context"
	"net/http"
	"time"
)

func main() {
	serverCtx, serverCancel := context.WithTimeout(context.Background(), 5*time.Second)
	minersCtx, minersCancel := context.WithCancel(context.Background())
	defer serverCancel()
	defer minersCancel()
	timeStart := time.Now()
	router := http.NewServeMux()
	eventBus := event.NewEventBus()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	//Services:
	minerService := miner.NewMinerService(miner.MinerServiceDeps{
		EventBus: eventBus,
	})

	walletService := wallet.NewWalletService(wallet.ServiceWalletDeps{
		EventBus: eventBus,
	},
	)

	statsService := stats.NewStatsMinersService(stats.StatsMinersServiceDeps{
		EventBus: eventBus,
	},
	)

	equipmentService := equipment.NewEquipmentService()

	StartGame(eventBus, minersCtx)

	//Handlers:

	handler.NewMinersHandler(
		router,
		handler.MinersHandlerDeps{
			Ctx:             minersCtx,
			MinerManager:    minerService,
			BallanceManager: walletService,
		},
	)

	handler.NewEquipmentHandler(
		router,
		handler.EquipmentHandlerDeps{
			EquipmentManager: equipmentService,
			BallanceManager:  walletService,
		},
	)

	handler.NewStatsHandler(
		router,
		handler.StatsHandlerDeps{
			StatsProvider: statsService,
		},
	)

	handler.NewEnterpriseHandler(
		router,
		handler.EnterpriseHandlerDeps{
			EnterpriseManager: equipmentService,
			BalanceManager:    walletService,
			StatsProvider:     statsService,
			MinersCancel:      minersCancel,
			Ctx:               serverCtx,
			Server:            &server,
			TimeStart:         &timeStart,
		},
	)

	server.ListenAndServe()

}

func StartGame(e *event.EventBus, ctx context.Context) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				e.Publish(event.Event{
					Type: event.EventMinerMined,
					Data: types.Coal(1),
				})
			}
		}

	}()
}
