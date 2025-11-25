package app

import (
	"MinersGame/internal/domain/equipment"
	"MinersGame/internal/domain/miner"
	"MinersGame/internal/domain/stats"
	"MinersGame/internal/domain/wallet"
	"MinersGame/pkg/event"
	"context"
	"net/http"
	"time"
)

type App struct {
	Server   *http.Server
	EventBus *event.EventBus

	MinerService     *miner.MinerService
	WalletService    *wallet.WalletService
	StatsService     *stats.StatsService
	EquipmentService *equipment.EquipmentService

	AppCtx context.Context
	Cancel context.CancelFunc

	TimeStart *time.Time
}

func New(ctx context.Context) App {
	timeStart := time.Now()
	appCtx, cancel := context.WithCancel(ctx)

	eventBus := event.NewEventBus()

	//Services:

	minerService := miner.NewMinerService(miner.MinerServiceDeps{
		EventBus: eventBus,
	})

	walletService := wallet.NewWalletService(wallet.ServiceWalletDeps{
		EventBus: eventBus,
	},
	)

	statsService := stats.NewStatsMinersService(stats.StatsServiceDeps{
		EventBus: eventBus,
	},
	)

	equipmentService := equipment.NewEquipmentService()

	return App{
		Server:   &http.Server{},
		EventBus: eventBus,

		MinerService:     minerService,
		WalletService:    walletService,
		StatsService:     statsService,
		EquipmentService: equipmentService,

		AppCtx: appCtx,
		Cancel: cancel,

		TimeStart: &timeStart,
	}
}
